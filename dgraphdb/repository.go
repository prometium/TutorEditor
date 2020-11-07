package dgraphdb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	"editorsvc"
	"editorsvc/utils"
)

type repository struct {
	db *dgo.Dgraph
}

// New returns a repository backed by Dgraph
func New(db *dgo.Dgraph) editorsvc.Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Setup(ctx context.Context) error {
	err := repo.db.Alter(ctx, &api.Operation{
		Schema:          schema,
		RunInBackground: true,
	})
	return err
}

func (repo *repository) AddScript(ctx context.Context, name string, frames []editorsvc.Frame) (string, error) {
	script := editorsvc.Script{
		UID:    "_:script",
		Name:   name,
		Frames: frames,
		DType:  []string{"Script"},
	}
	for i := range script.Frames {
		frame := &script.Frames[i]
		frame.UID = fmt.Sprintf("_:frame-%s", frame.UID)
		frame.Task.UID = fmt.Sprintf("_:task-%d", utils.Hash(frame.Task.Text))
		frame.Hint.UID = fmt.Sprintf("_:hint-%d", utils.Hash(frame.Hint.Text))
		frame.DType = []string{"Frame"}
		for j := range frame.Actions {
			action := &frame.Actions[j]
			action.NextFrame.UID = fmt.Sprintf("_:frame-%s", action.NextFrame.UID)
			action.DType = []string{"Action"}
		}
	}

	scriptB, err := json.Marshal(script)
	if err != nil {
		return "", err
	}

	mu := &api.Mutation{
		SetJson:   scriptB,
		CommitNow: true,
	}

	assigned, err := repo.db.NewTxn().Mutate(ctx, mu)
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
	resp, err := repo.db.NewTxn().Query(ctx, q)
	if err != nil {
		return nil, err
	}

	var decode struct {
		Scripts []editorsvc.Script
	}
	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		return nil, err
	}
	return decode.Scripts, nil
}
