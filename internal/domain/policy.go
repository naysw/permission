package domain

import "context"

type Principal struct {
	ID   string
	Type string
}

type Policy struct {
	ID          string
	Name        string
	Description *string
	Document    string
}

type GetListPolicy struct {
	Skip          int
	Limit         int
	IDs           []string
	Name          *string
	PrincipalID   *string
	PrincipalType *string
}

type UpdatePolicy struct {
	Name        string
	Description string
	Document    string
}

type AttachPolicyInput struct {
	PrincipalID   string
	PrincipalType string
	PolicyIDs     []string
}

type DetachPolicyInput struct {
	PrincipalID   string
	PrincipalType string
	PolicyIDs     []string
}

type PolicyRepo interface {
	GetList(ctx context.Context, input GetListPolicy) ([]Policy, error)
	Create(ctx context.Context, p Policy) error
	Update(ctx context.Context, id string, p UpdatePolicy) error
	Delete(ctx context.Context, id string) error
	Attach(ctx context.Context, input AttachPolicyInput) error
	Detach(ctx context.Context, input DetachPolicyInput) error
}
