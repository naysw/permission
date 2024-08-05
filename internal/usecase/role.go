package usecase

import (
	"context"

	"github.com/naysw/permission/internal/domain"
)

type RoleUsecase struct {
	roleRepo   domain.RoleRepo
	policyRepo domain.PolicyRepo
}

func NewRoleUsecase(roleRepo domain.RoleRepo, policyRepo domain.PolicyRepo) *RoleUsecase {
	return &RoleUsecase{
		roleRepo:   roleRepo,
		policyRepo: policyRepo,
	}
}

type CreateRoleInput struct {
	Name        string
	Description *string
}

func (r RoleUsecase) CreateRole(ctx context.Context, ip *CreateRoleInput) (string, error) {
	if ip == nil {
		return "", domain.ErrInvalidInput
	}

	panic("not implemented")
}
