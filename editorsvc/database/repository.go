package dgraphdb

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
	script = prepareScriptForStorage(script)

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

func (repo *repository) GetScript(ctx context.Context, id string) (*editorsvc.Script, error) {
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
	if len(decode.Script) == 0 {
		return nil, nil
	}
	for i := range decode.Script[0].Frames {
		frame := &decode.Script[0].Frames[i]

		for j := range frame.Actions {
			action := &frame.Actions[j]

			if action.SwitchPicturesJSON == "" {
				continue
			}

			err := json.Unmarshal([]byte(action.SwitchPicturesJSON), &action.SwitchPictures)
			if err != nil {
				return nil, err
			}
			action.SwitchPicturesJSON = ""
		}
	}
	return &decode.Script[0], nil
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

func (repo *repository) UpdateScript(
	ctx context.Context, script *editorsvc.Script, frameIdsToDel []string, actionIdsToDel []string,
) (map[string]string, error) {
	q := `query script($scriptId: string, $frameIdsToDel: string, $actionIdsToDel: string) {
		script as var(func: uid($scriptId)) @filter(type("Script"))
		framesToDel as var(func: uid($frameIdsToDel))
		actionsToDel as var(func: uid($actionIdsToDel))
		var(func: uid(actionsToDel)) {
			~actions {
				prevFrames as uid
			}
		}
	}`

	mu1 := &api.Mutation{
		DelNquads: []byte(`
			uid(script) <frames> uid(framesToDel) .
			uid(prevFrames) <actions> uid(actionsToDel) .
			uid(framesToDel) * * .
			uid(actionsToDel) * * .
		`),
	}

	script = prepareScriptForStorage(script)
	scriptB, err := json.Marshal(script)
	if err != nil {
		return nil, err
	}

	mu2 := &api.Mutation{
		SetJson: scriptB,
	}

	req := &api.Request{
		Query:     q,
		Mutations: []*api.Mutation{mu1, mu2},
		Vars: map[string]string{
			"$scriptId":       script.UID,
			"$frameIdsToDel":  fmt.Sprintf("[%s]", strings.Join(frameIdsToDel, ",")),
			"$actionIdsToDel": fmt.Sprintf("[%s]", strings.Join(actionIdsToDel, ",")),
		},
		CommitNow: true,
	}

	assigned, err := repo.dg.NewTxn().Do(ctx, req)
	if err != nil {
		return nil, err
	}
	return assigned.Uids, nil
}

func prepareScriptForStorage(script *editorsvc.Script) *editorsvc.Script {
	script.DType = []string{"Script"}

	for i := range script.Frames {
		frame := &script.Frames[i]
		script.Frames[i].DType = []string{"Frame"}

		for j := range frame.Actions {
			action := &frame.Actions[j]
			action.DType = []string{"Action"}

			if action.SwitchPictures == nil {
				continue
			}

			b, err := json.Marshal(action.SwitchPictures)
			if err != nil {
				action.SwitchPicturesJSON = ""
			}
			action.SwitchPicturesJSON = string(b)
			action.SwitchPictures = nil
		}
	}

	return script
}
