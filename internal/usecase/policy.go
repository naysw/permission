package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
	"github.com/naysw/permission/internal/domain"
)

type PolicyUsecase struct {
	policyRepo domain.PolicyRepo
}

type Policy struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Document    string  `json:"document"`
}

func NewPolicyUsecase(policyRepo domain.PolicyRepo) *PolicyUsecase {
	return &PolicyUsecase{
		policyRepo: policyRepo,
	}
}

type GetListPolicy struct {
	Skip  int
	Limit int
	IDs   []string
	Name  *string
}

type CreatePolicyInput struct {
	Name        string
	Description *string
	Document    string
}

func (u PolicyUsecase) GetList(ctx context.Context, input *GetListPolicy) ([]Policy, error) {
	dm := domain.GetListPolicy{
		Limit: 20,
	}
	if input != nil {
		if input.Name != nil {
			dm.Name = input.Name
		}
		if input.IDs != nil {
			dm.IDs = input.IDs
		}
		if input.Skip > 0 {
			dm.Skip = input.Skip
		}
		if input.Limit > 0 {
			if input.Limit > 100 {
				input.Limit = 100
			} else {
				dm.Limit = input.Limit
			}
		}
	}
	ps, err := u.policyRepo.GetList(ctx, dm)
	if err != nil {
		return nil, err
	}

	var policies []Policy
	for _, p := range ps {
		policies = append(policies, Policy{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Document:    p.Document,
		})
	}

	return policies, nil
}

func (u PolicyUsecase) CreatePolicy(
	ctx context.Context,
	input CreatePolicyInput) (*string, error) {

	// TODO:
	// parse input.Document to check if it is a valid document
	id := uuid.New().String()
	if err := u.policyRepo.Create(
		ctx,
		domain.Policy{
			ID:          id,
			Name:        input.Name,
			Description: input.Description,
			Document:    input.Document,
		},
	); err != nil {
		return nil, err
	}

	return &id, nil
}

func (u PolicyUsecase) Authorized(
	ctx context.Context,
	entities cedar.Entities,
	req cedar.Request) (bool, error) {
	// get policies by principal

	var policyDocs []string
	skip := 0
	limit := 100

	for {
		ps, err := u.policyRepo.GetList(
			ctx,
			domain.GetListPolicy{
				Skip:          skip,
				Limit:         limit,
				PrincipalID:   &req.Principal.ID,
				PrincipalType: &req.Principal.Type,
			},
		)
		if err != nil {
			return false, err
		}

		for _, p := range ps {
			fmt.Println("document", p.Document)
			policyDocs = append(policyDocs, p.Document)
		}
		if len(ps) < limit {
			break
		}
		skip += limit
	}
	pb := []byte(strings.Join(policyDocs, "\n"))
	ps, err := cedar.NewPolicySet("policies.cedar", pb)
	if err != nil {
		return false, err
	}

	ok, _ := ps.IsAuthorized(entities, req)
	if ok == cedar.Allow {
		return true, nil
	}

	return false, nil
}

type AttachPolicyInput struct {
	PrincipalID   string
	PrincipalType string
	PolicyIDs     []string
}

func (u PolicyUsecase) AttachPolicy(
	ctx context.Context,
	input AttachPolicyInput) error {

	return u.policyRepo.Attach(
		ctx,
		domain.AttachPolicyInput{
			PrincipalID:   input.PrincipalID,
			PrincipalType: input.PrincipalType,
			PolicyIDs:     input.PolicyIDs,
		},
	)
}
