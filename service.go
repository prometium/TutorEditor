package editorsvc

import (
	"context"
	"io"
)

// Service provides some "date capabilities" to application
type Service interface {
	AddRawScript(ctx context.Context, name string, archiveReader io.Reader) (string, error)
}
