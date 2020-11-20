package dgraphdb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"

	"github.com/prometium/tutoreditor/editorsvc"
)

type repository struct {
	dg *dgo.Dgraph
}

// New returns a repository backed by Dgraph
func New(dg *dgo.Dgraph) editorsvc.Repository {
	return &repository{
		dg: dg,
	}
}

func (repo *repository) Setup(ctx context.Context) error {
	err := repo.dg.Alter(ctx, &api.Operation{
		Schema:          schema,
		RunInBackground: true,
	})
	return err
}

func (repo *repository) AddScript(ctx context.Context, script *editorsvc.Script) (string, error) {
	script = classifyScript(script)

	script.UID = "_:script"
	for i := range script.Frames {
		frame := &script.Frames[i]
		frame.UID = fmt.Sprintf("_:f%s", frame.UID)

		for j := range frame.Actions {
			action := &frame.Actions[j]
			action.UID = ""

			nextFrame := action.NextFrame
			nextFrame.UID = fmt.Sprintf("_:f%s", nextFrame.UID)
		}
	}
	script.FirstFrame.UID = fmt.Sprintf("_:f%s", script.FirstFrame.UID)

	scriptB, err := json.Marshal(script)
	if err != nil {
		return "", err
	}

	mu := &api.Mutation{
		SetJson:   scriptB,
		CommitNow: true,
	}

	assigned, err := repo.dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return "", err
	}
	return assigned.Uids["script"], nil
}

func (repo *repository) GetScriptsList(ctx context.Context) ([]editorsvc.Script, error) {
	q := `{
		scripts(func: type("Script")) {
			uid
			name
			version
		}
	}`

	res, err := repo.dg.NewTxn().Query(ctx, q)
	if err != nil {
		return nil, err
	}

	var decode struct {
		Scripts []editorsvc.Script
	}
	if err := json.Unmarshal(res.GetJson(), &decode); err != nil {
		return nil, err
	}
	return decode.Scripts, nil
}

func (repo *repository) GetScript(ctx context.Context, id string) ([]editorsvc.Script, error) {
	q := `query script($id: string) {
		script(func: uid($id)) @filter(type("Script")) {
			uid
			name
			version
			firstFrame {
			  uid
			}
			frames {
				uid
				expand(_all_) {
					uid
					expand(_all_) {
						uid
					}
				}
			}
		}
	}`

	res, err := repo.dg.NewTxn().QueryWithVars(ctx, q, map[string]string{"$id": id})
	if err != nil {
		return nil, err
	}

	var decode struct {
		Script []editorsvc.Script
	}
	if err := json.Unmarshal(res.GetJson(), &decode); err != nil {
		return nil, err
	}
	return decode.Script, nil
}

func (repo *repository) DeleteScript(ctx context.Context, id string) error {
	q := `query script($id: string) {
		script as var(func: uid($id)) @filter(type("Script")) {
			expand(_all_) {
				depth2 as uid
				expand(_all_) {
					depth3 as uid
				}
			}
		}
	}`

	mu := &api.Mutation{
		DelNquads: []byte(`
			uid(script) * * .
			uid(depth2) * * .
			uid(depth3) * * .
		`),
	}

	req := &api.Request{
		Query:     q,
		Mutations: []*api.Mutation{mu},
		Vars:      map[string]string{"$id": id},
		CommitNow: true,
	}

	if _, err := repo.dg.NewTxn().Do(ctx, req); err != nil {
		return err
	}
	return nil
}

func (repo *repository) UpdateScript(ctx context.Context, script *editorsvc.Script) (map[string]string, error) {
	script = classifyScript(script)
	scriptB, err := json.Marshal(script)
	if err != nil {
		return nil, err
	}

	mu := &api.Mutation{
		SetJson: scriptB,
	}

	req := &api.Request{
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	assigned, err := repo.dg.NewTxn().Do(ctx, req)
	if err != nil {
		return nil, err
	}
	return assigned.Uids, nil
}

func (repo *repository) AddBranch(ctx context.Context, script *editorsvc.Script, branch *editorsvc.Branch) (map[string]string, error) {
	if len(branch.ConnectedFrames) == 0 {
		return nil, nil
	}

	frames := branch.ConnectedFrames

	for i := range frames {
		frame := &frames[i]
		frame.UID = fmt.Sprintf("_:%s", frame.UID)

		action := &frame.Actions[0]
		action.UID = fmt.Sprintf("_:%s", action.UID)
		action.NextFrame.UID = fmt.Sprintf("_:%s", action.NextFrame.UID)
	}

	firstMainFrame := editorsvc.Frame{
		UID:     branch.FirstMainFrameID,
		Actions: frames[0].Actions,
	}

	frames[len(frames)-1].Actions[0].NextFrame.UID = branch.LastMainFrameID

	frames = append(frames[1:], firstMainFrame)

	script.Frames = frames
	return repo.UpdateScript(ctx, script)
}

func (repo *repository) DeleteBranch(ctx context.Context, script *editorsvc.Script, branchToDelete *editorsvc.BranchToDelete) error {
	q := `query script($scriptId: string, $branchFrameId: string, $firstActionId: string, $lastActionId: string) {
		var(func: uid($scriptId)) @filter(type("Script")) {
			frames as frames {
				actions as actions
			}
		}
		root as var(func: uid($branchFrameId)) @filter(uid(frames))
		first as var(func: uid($firstActionId)) @filter(uid(actions))
		last as var(func: uid($lastActionId)) @filter(uid(actions))
		path as shortest(from: uid(first), to: uid(last)) {
			actions
			nextFrame
		}
	}`

	mu := &api.Mutation{
		DelNquads: []byte(`
			uid(root) <actions> uid(first) .
			uid(path) * * .
		`),
	}

	req := &api.Request{
		Query:     q,
		Mutations: []*api.Mutation{mu},
		Vars: map[string]string{
			"$scriptId":      script.UID,
			"$branchFrameId": branchToDelete.BranchFrameID,
			"$firstActionId": branchToDelete.FirstActionID,
			"$lastActionId":  branchToDelete.LastActionID,
		},
		CommitNow: true,
	}

	if _, err := repo.dg.NewTxn().Do(ctx, req); err != nil {
		return err
	}
	if _, err := repo.UpdateScript(ctx, script); err != nil {
		return err
	}
	return nil
}

func (repo *repository) DeleteFrame(ctx context.Context, script *editorsvc.Script, id string) error {
	q := `query frame($scriptId: string, $frameId: string) {
		var(func: uid($scriptId)) @filter(type("Script")) {
			frames as frames
		}
		frame as var(func: uid($frameId)) @filter(uid(frames)) {
			prevAction: ~nextFrame {
				prevAction as uid
			}
			actions {
				nextFrame {
					nextFrame as uid
				}
			}
		}
	}`

	mu := &api.Mutation{
		DelNquads: []byte(`
			uid(frame) * * .
			uid(prevAction) <nextFrame> uid(frame) .
		`),
		SetNquads: []byte(`
			uid(prevAction) <nextFrame> uid(nextFrame) .
		`),
	}

	req := &api.Request{
		Query:     q,
		Mutations: []*api.Mutation{mu},
		Vars: map[string]string{
			"$scriptId": script.UID,
			"$frameId":  id,
		},
		CommitNow: true,
	}

	if _, err := repo.dg.NewTxn().Do(ctx, req); err != nil {
		return err
	}
	if _, err := repo.UpdateScript(ctx, script); err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetScriptVersion(ctx context.Context, id string) (string, error) {
	q := `query script($id: string) {
		script(func: uid($id)) @filter(type("Script")) {
			version
		}
	}`

	res, err := repo.dg.NewTxn().QueryWithVars(ctx, q, map[string]string{"$id": id})
	if err != nil {
		return "", err
	}

	var decode struct {
		Script []editorsvc.Script
	}
	if err := json.Unmarshal(res.GetJson(), &decode); err != nil {
		return "", err
	} else if len(decode.Script) > 0 {
		return decode.Script[0].Version, nil
	}
	return "", nil
}

func classifyScript(script *editorsvc.Script) *editorsvc.Script {
	script.DType = []string{"Script"}
	script.Frames = classifyFrames(script.Frames)
	return script
}

func classifyFrames(frames []editorsvc.Frame) []editorsvc.Frame {
	for i := range frames {
		frame := &frames[i]
		frames[i].DType = []string{"Frame"}
		for j := range frame.Actions {
			frame.Actions[j].DType = []string{"Action"}
		}
	}
	return frames
}
