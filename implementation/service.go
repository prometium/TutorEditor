package implementation

import (
	"context"
	"fmt"
	"io"
	"os"

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

func (s *service) Setup(ctx context.Context) error {
	return s.repository.Setup(ctx)
}

func (s *service) AddRawScript(ctx context.Context, name string, fileReader io.Reader) (string, error) {
	var rs rawScript
	if err := rs.init(fileReader); err != nil {
		return "", err
	}

	imagesDir := fmt.Sprintf("assets/images/")
	os.MkdirAll(imagesDir, os.ModePerm)

	linksMap, err := rs.storeImages(imagesDir)
	if err != nil {
		return "", err
	}

	frames, err := rs.generateFrames(imagesDir, linksMap)
	if err != nil {
		return "", err
	}

	id, err := s.repository.AddScript(ctx, name, frames)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *service) GetScriptsList(ctx context.Context) ([]editorsvc.Script, error) {
	return s.repository.GetScriptsList(ctx)
}
