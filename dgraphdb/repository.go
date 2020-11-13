package dgraphdb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	"editorsvc"
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
	script = configureScript(script)
	script.UID = "_:script"

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
		scripts(func: eq(dgraph.type, "Script")) {
			uid
			name
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
		script(func: uid($id)) @filter(eq(dgraph.type, "Script")) {
			uid
			name
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
		script as var(func: uid($id)) {
			expand(_all_) {
				depth2 as uid
			}
		}
	}`

	mu := &api.Mutation{
		DelNquads: []byte(`
			uid(script) * * .
			uid(depth2) * * .
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

func (repo *repository) UpdateScript(ctx context.Context, script *editorsvc.Script) error {
	q := `query script($id: string) {
		script as var(func: uid($id)) {
			expand(_all_) {
				depth2 as uid
			}
		}
	}`

	mu1 := &api.Mutation{
		DelNquads: []byte(`
			uid(script) <frames> * .
			uid(depth2) * * .
		`),
	}

	script = configureScript(script)
	scriptB, err := json.Marshal(script)
	if err != nil {
		return err
	}

	mu2 := &api.Mutation{
		SetJson: scriptB,
	}

	req := &api.Request{
		Query:     q,
		Mutations: []*api.Mutation{mu1, mu2},
		Vars:      map[string]string{"$id": script.UID},
		CommitNow: true,
	}

	if _, err := repo.dg.NewTxn().Do(ctx, req); err != nil {
		return err
	}

	return nil
}

func configureScript(script *editorsvc.Script) *editorsvc.Script {
	script.DType = []string{"Script"}

	for i := range script.Frames {
		frame := &script.Frames[i]
		frame.UID = fmt.Sprintf("_:frame-%s", frame.UID)
		frame.DType = []string{"Frame"}

		for j := range frame.Actions {
			action := &frame.Actions[j]
			action.DType = []string{"Action"}

			action.NextFrame.UID = fmt.Sprintf("_:frame-%s", action.NextFrame.UID)
		}
	}

	script.FirstFrame.UID = fmt.Sprintf("_:frame-%s", script.FirstFrame.UID)

	return script
}
