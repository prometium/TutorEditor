package implementation

import (
	"context"
	"io"

	"editorsvc"
)

type service struct {
	repository editorsvc.Repository
}

// NewService makes a new Service
func NewService(rep editorsvc.Repository) editorsvc.Service {
	return &service{
		repository: rep,
	}
}

func (s *service) AddRawScript(ctx context.Context, name string, fileReader io.ReadCloser) (string, error) {
	defer fileReader.Close()

	var rs rawScript
	if err := rs.init(fileReader); err != nil {
		return "", err
	}

	linksMap, err := rs.saveImages(ctx, "assets/images/")
	if err != nil {
		return "", err
	}

	script, err := rs.createScript(name, linksMap)
	if err != nil {
		return "", err
	}

	id, err := s.repository.AddScript(ctx, script)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *service) GetScriptsList(ctx context.Context) ([]editorsvc.Script, error) {
	return s.repository.GetScriptsList(ctx)
}

func (s *service) GetScript(ctx context.Context, id string) (*editorsvc.Script, error) {
	script, err := s.repository.GetScript(ctx, id)
	if err != nil {
		return nil, err
	} else if len(script) == 0 {
		return nil, editorsvc.ErrScriptNotFound
	}
	return &script[0], err
}

func (s *service) DeleteScript(ctx context.Context, id string) error {
	return s.repository.DeleteScript(ctx, id)
}

func (s *service) UpdateScript(ctx context.Context, script *editorsvc.Script) error {
	return s.repository.UpdateScript(ctx, script)
}

func (s *service) AddBranch(ctx context.Context, branch *editorsvc.Branch) (map[string]string, error) {
	return s.repository.AddBranch(ctx, branch)
}

func (s *service) DeleteBranch(ctx context.Context, branchToDelete *editorsvc.BranchToDelete) error {
	return s.repository.DeleteBranch(ctx, branchToDelete)
}
