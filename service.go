package editorsvc

import (
	"context"
	"io"
)

var ()

// Service provides some "date capabilities" to application
type Service interface {
	Setup(ctx context.Context) error
	AddRawScript(ctx context.Context, name string, archiveReader io.Reader) (string, error)
	GetScriptsList(ctx context.Context) ([]Script, error)
}
