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

func (s *service) GetScript(ctx context.Context, id string) ([]editorsvc.Script, error) {
	return s.repository.GetScript(ctx, id)
}
