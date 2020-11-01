package implementation

import (
	"context"

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

func (service) Status(ctx context.Context) (string, error) {
	return "Ok", nil
}

func (service) TransformScript(ctx context.Context, script string) (int, error) {
	return 0, nil
}
