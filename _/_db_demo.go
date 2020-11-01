package _db_demo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

func db() {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	op := &api.Operation{}
	op.Schema = `
		type Frame {
			actions
			task
			hint
		}

		type Action {
			nextFrame
		}

		type Task {
			taskText
		}
		  
		type Hint {
			hintText
		}

		actions: [uid] .
		nextFrame: uid .
		task: uid .
		hint: uid .
		taskText: string .
		hintText: string .
	`

	ctx := context.Background()
	err = dgraphClient.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	type NextFrame struct {
		UID string `json:"uid,omitempty"`
	}

	type Action struct {
		NextFrame NextFrame `json:"next_frame,omitempty"`
		DType     []string  `json:"dgraph.type,omitempty"`
	}

	type Frame struct {
		UID     string   `json:"uid,omitempty"`
		Actions []Action `json:"actions,omitempty"`
		DType   []string `json:"dgraph.type,omitempty"`
	}

	mutations := make([]*api.Mutation, 2)
	for i, x := range []int{1, 2} {
		frame := Frame{
			UID: fmt.Sprintf("_:frame%d", x),
			DType: []string{"Frame"},
		}

		frameB, _ := json.Marshal(frame)
		if err != nil {
			log.Fatal(err)
		}

		mu := &api.Mutation{
			SetJson: frameB,
		}

		mutations[i] = mu
	}

	req := &api.Request{CommitNow: true, Mutations: mutations}
	assigned, err := dgraphClient.NewTxn().Do(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(assigned)

	updatedFrame := Frame{
		UID: assigned.Uids["frame1"],
		Actions: []Action{
			Action{
				NextFrame: NextFrame{
					UID: assigned.Uids["frame2"],
				},
				DType:     []string{"Action"},
			}},
	}

	updatedFrameB, _ := json.Marshal(updatedFrame)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		SetJson:   updatedFrameB,
		CommitNow: true,
	}

	assigned, err = dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(assigned)
}
