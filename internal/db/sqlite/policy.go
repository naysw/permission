package sqlite

import (
	"context"
	"time"

	"github.com/naysw/permission/internal/domain"
	"gorm.io/gorm"
)

type Policy struct {
	ID          string         `gorm:"primaryKey"`
	Name        string         `gorm:"column:name;NOT NULL;VARCHAR(255)"`
	Description *string        `gorm:"column:description"`
	Document    string         `gorm:"column:document;NOT NULL"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

type PrincipalHasPolicy struct {
	PrincipalID   string         `gorm:"column:principal_id;NOT NULL;index:idx_principal_id_and_policy_id,unique,where:deleted_at IS NULL"`
	PrincipalType string         `gorm:"column:principal_type;NOT NULL;index:idx_principal_id_and_policy_id,unique,where:deleted_at IS NULL"`
	PolicyID      string         `gorm:"column:policy_id;NOT NULL;index:idx_principal_id_and_policy_id,unique,where:deleted_at IS NULL"`
	Policy        Policy         `gorm:"foreignKey:PolicyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (PrincipalHasPolicy) TableName() string {
	return "principal_has_policies"
}

func (Policy) TableName() string {
	return "policies"
}

func (p Policy) ToDomain() domain.Policy {
	return domain.Policy{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Document:    p.Document,
	}
}

type PolicyRepo struct {
	db *gorm.DB
}

func (r PolicyRepo) GetList(ctx context.Context, input domain.GetListPolicy) ([]domain.Policy, error) {
	var ps []Policy
	db := r.db.WithContext(ctx).Model(&Policy{})
	if input.Name != nil && len(*input.Name) > 0 {
		db.Where("LOWER(name) LIKE ?", "%"+*input.Name+"%")
	}
	if input.IDs != nil && len(input.IDs) > 0 {
		db.Where("id IN ?", input.IDs)
	}
	if input.Skip > 0 {
		db.Offset(input.Skip)
	}
	if input.Limit > 0 {
		db.Limit(input.Limit)
	}

	if input.PrincipalID != nil && input.PrincipalType != nil {
		db = db.Joins(
			`JOIN principal_has_policies
				ON policies.id = principal_has_policies.policy_id`,
		).Where(
			"principal_has_policies.principal_id = ? AND principal_has_policies.principal_type = ?",
			*input.PrincipalID,
			*input.PrincipalType,
		).Where("principal_has_policies.deleted_at IS NULL")
	}

	if err := db.Find(&ps).Error; err != nil {
		return nil, err
	}

	var policies []domain.Policy
	for _, p := range ps {
		policies = append(policies, p.ToDomain())
	}

	return policies, nil
}

func (r PolicyRepo) Create(
	ctx context.Context,
	p domain.Policy) error {

	po := Policy{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Document:    p.Document,
	}

	return r.db.WithContext(ctx).Create(&po).Error
}

func (r PolicyRepo) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (r PolicyRepo) Update(ctx context.Context, id string, p domain.UpdatePolicy) error {
	panic("unimplemented")
}

func (r PolicyRepo) Attach(ctx context.Context, input domain.AttachPolicyInput) error {
	var phps []PrincipalHasPolicy
	for _, pid := range input.PolicyIDs {
		phps = append(phps, PrincipalHasPolicy{
			PrincipalID:   input.PrincipalID,
			PrincipalType: input.PrincipalType,
			PolicyID:      pid,
		})
	}

	return r.db.WithContext(ctx).Create(&phps).Error
}

func (r PolicyRepo) Detach(ctx context.Context, input domain.DetachPolicyInput) error {
	return r.db.WithContext(ctx).
		Where(`
      principal_id = ? AND 
      principal_type = ? AND policy_id IN ?
      `,
			input.PrincipalID,
			input.PrincipalType,
			input.PolicyIDs,
		).
		Delete(&PrincipalHasPolicy{}).Error
}

func NewPolicyRepo(db *gorm.DB) domain.PolicyRepo {
	return &PolicyRepo{db: db}
}
