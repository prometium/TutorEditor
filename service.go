package editorsvc

import (
	"context"
	"errors"
	"io"
)

var (
	// ErrScriptAlreadyExists is returned when script with this name already exists
	ErrScriptAlreadyExists = errors.New("script already exists")
)

// Service provides some "date capabilities" to application
type Service interface {
	Setup(ctx context.Context) error
	AddRawScript(ctx context.Context, name string, archiveReader io.Reader) (string, error)
}
