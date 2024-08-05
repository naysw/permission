package sqlite

import (
	"context"
	"time"

	"github.com/naysw/permission/internal/domain"
	"gorm.io/gorm"
)

type Role struct {
	ID          string         `gorm:"primaryKey"`
	Name        string         `gorm:"column:name;NOT NULL;VARCHAR(255)"`
	Description *string        `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Role) TableName() string {
	return "roles"
}

type RoleRepo struct {
	db *gorm.DB
}

func (*RoleRepo) Create(ctx context.Context, r domain.Role) (string, error) {
	panic("unimplemented")
}

func (r *RoleRepo) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (*RoleRepo) Update(ctx context.Context, id string, r domain.UpdateRole) error {
	panic("unimplemented")
}

func NewRoleRepo(db *gorm.DB) domain.RoleRepo {
	return &RoleRepo{db: db}
}
