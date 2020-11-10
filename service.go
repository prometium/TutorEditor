package editorsvc

import (
	"context"
	"io"
)

// Service provides some "date capabilities" to application
type Service interface {
	AddRawScript(ctx context.Context, name string, archiveReader io.ReadCloser) (string, error)
	GetScriptsList(ctx context.Context) ([]Script, error)
	GetScript(ctx context.Context, id string) ([]Script, error)
}
