package editorsvc

import (
	"context"
	"errors"
	"io"
)

var (
	// ErrScriptNotFound denotes a script was not found
	ErrScriptNotFound = errors.New("script not found")
)

// Service provides some "date capabilities" to application
type Service interface {
	AddRawScript(ctx context.Context, name string, archiveReader io.ReadCloser) (string, error)
	GetScriptsList(ctx context.Context) ([]Script, error)
	GetScript(ctx context.Context, id string) (*Script, error)
	DeleteScript(ctx context.Context, id string) error
	UpdateScript(ctx context.Context, script *Script) error
}
