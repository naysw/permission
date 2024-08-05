package domain

import "context"

type Role struct {
	ID   string
	Name string
}

type UpdateRole struct {
	Name string
}

type RoleRepo interface {
	Create(ctx context.Context, r Role) (string, error)
	Update(ctx context.Context, id string, r UpdateRole) error
	Delete(ctx context.Context, id string) error
}
