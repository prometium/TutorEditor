package dgraphdb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	"editorsvc"
)

type repository struct {
	db *dgo.Dgraph
}

// New returns a repository backed by Dgraph
func New(db *dgo.Dgraph) (editorsvc.Repository, error) {
	return &repository{
		db: db,
	}, nil
}

func (repo *repository) AddScript(ctx context.Context, frames []editorsvc.Frame) error {
	mutations := make([]*api.Mutation, len(frames))
	for i, frame := range frames {
		frameB, err := json.Marshal(frame)
		if err != nil {
			log.Fatal(err)
		}

		mu := &api.Mutation{
			SetJson: frameB,
		}
		mutations[i] = mu
	}

	fmt.Println("HERE")

	req := &api.Request{CommitNow: true, Mutations: mutations}
	assigned, err := repo.db.NewTxn().Do(ctx, req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(assigned)

	return nil
}
