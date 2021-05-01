package editorsvc

import (
	"context"
	"errors"
	"io"
)

var (
	// ErrFileNotAttached denotes the file was not attached
	ErrFileNotAttached = errors.New("file not attached")
	// ErrScriptNotFound denotes the script was not found
	ErrScriptNotFound = errors.New("script not found")
	// ErrVersionsDoNotMatch denotes the script versions do not match
	ErrVersionsDoNotMatch = errors.New("script versions do not match")
)

// Service provides some "date capabilities" to application
type Service interface {
	AddScriptArchive(ctx context.Context, name string, archiveReader io.ReadCloser) (string, error)
	GetScriptArchive(ctx context.Context, id string) ([]byte, error)
	GetScriptsList(ctx context.Context) ([]Script, error)
	GetScript(ctx context.Context, id string) (*Script, error)
	DeleteScript(ctx context.Context, id string) error
	UpdateScript(ctx context.Context, script *Script, frameIdsToDel []string, actionIdsToDel []string) (map[string]string, error)
	CopyScript(ctx context.Context, script *Script) (string, error)
}
