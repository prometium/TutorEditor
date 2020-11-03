package editorsvc

import (
	"context"
	"io"
)

// Service provides some "date capabilities" to application
type Service interface {
	Status(ctx context.Context) (string, error)
	AddRawScript(ctx context.Context, archiveReader io.Reader) (int, error)
}
