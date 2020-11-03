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
		UID:    "script",
		Name:   name,
		Frames: frames,
	}
	prepareScript(&script)

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

func prepareScript(script *editorsvc.Script) {
	script.UID = fmt.Sprintf("_:%s", script.UID)
	script.DType = []string{"Script"}
	for i := range script.Frames {
		frame := &script.Frames[i]
		frame.UID = fmt.Sprintf("_:%s", frame.UID)
		frame.DType = []string{"Frame"}
		for j := range frame.Actions {
			action := &frame.Actions[j]
			action.NextFrame.UID = fmt.Sprintf("_:%s", action.NextFrame.UID)
			action.DType = []string{"Action"}
		}
	}
}
