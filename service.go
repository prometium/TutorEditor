package editorsvc

import "context"

// Service provides some "date capabilities" to application
type Service interface {
	Status(ctx context.Context) (string, error)
	TransformScript(ctx context.Context, script string) (int, error)
}

// Repository describes the persistence on editor model
type Repository interface {
}
