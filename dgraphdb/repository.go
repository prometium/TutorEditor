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

func (repo *repository) AddScript(ctx context.Context, script editorsvc.Script) (string, error) {
	script.UID = "_:script"
	script.DType = []string{"Script"}
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
	script.FirstFrame.UID = script.Frames[0].UID

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
	res, err := repo.db.NewTxn().Query(ctx, q)
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
		script(func: eq(dgraph.type, "Script")) @filter(uid($id)) {
			uid
			name
			firstFrame {
				uid
			}
			frames {
				uid
				pictureLink
				actions {
					nextFrame {
				  		uid
					}
					actionType
					xLeft
					xRight
					yLeft
					yRight
					startXLeft
					startYLeft
					startXRight
					startYRight
					finishXLeft
					finishYLeft
					finishXRight
					finishYRight
					ticksCount
					key
					modKey
			  	}
			  	task {
					text
			  	}
			  	hint {
					text
			  	}
			}
		}
	}`

	res, err := repo.db.NewTxn().QueryWithVars(ctx, q, map[string]string{"$id": id})
	if err != nil {
		return nil, err
	}

	var decode struct {
		Script []editorsvc.Script
	}
	if err := json.Unmarshal(res.GetJson(), &decode); err != nil {
		return decode.Script, err
	}
	return decode.Script, nil
}
