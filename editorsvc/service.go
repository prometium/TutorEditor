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
	// ErrInvalidRequestParameters denotes the request parameters are invalid
	ErrInvalidRequestParameters = errors.New("invalid request parameters")
)

// Service provides some "date capabilities" to application
type Service interface {
	AddRawScript(ctx context.Context, name string, archiveReader io.ReadCloser) (string, error)
	GetScriptsList(ctx context.Context) ([]Script, error)
	GetScript(ctx context.Context, id string) (*Script, error)
	DeleteScript(ctx context.Context, id string) error
	UpdateScript(ctx context.Context, id string, script *Script) (string, map[string]string, error)
	CopyScript(ctx context.Context, script *Script) (string, error)
	AddBranch(ctx context.Context, script *Script, branch *Branch) (string, map[string]string, error)
	DeleteBranch(ctx context.Context, script *Script, branchToDelete *BranchToDelete) (string, error)
	AddFrame(ctx context.Context, script *Script, framesPair []Frame) (string, map[string]string, error)
	DeleteFrame(ctx context.Context, script *Script, id string) (string, error)
}
