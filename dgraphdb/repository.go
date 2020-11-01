package dgraphdb

import (
	"github.com/dgraph-io/dgo"

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
