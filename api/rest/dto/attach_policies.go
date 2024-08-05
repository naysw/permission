package dto

import (
	"strings"

	"github.com/naysw/permission/pkg/errs"
)

type AttachPolicy struct {
	PrincipalID   string   `json:"principal_id" validate:"required"`
	PrincipalType string   `json:"principal_type" validate:"required"`
	PolicyIDs     []string `json:"policy_ids" validate:"required"`
}

func (d AttachPolicy) Validate() (fes []errs.FieldErr) {
	em := strings.TrimSpace(d.PrincipalID)
	switch {
	case em == "":
		fes = append(fes, *errs.NewFieldErr("principal_id", "principal_id is required", "required"))
	case d.PrincipalType == "":
		fes = append(fes, *errs.NewFieldErr("principal_type", "principal_type is required", "required"))
	case len(d.PolicyIDs) == 0:
		fes = append(fes, *errs.NewFieldErr("policy_ids", "policy_ids is required", "required"))
	}

	return fes
}
