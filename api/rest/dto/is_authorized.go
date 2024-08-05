package dto

import (
	"github.com/cedar-policy/cedar-go"
	"github.com/naysw/permission/pkg/errs"
)

type AuthorizedDto struct {
	Request  cedar.Request  `json:"request" validate:"required"`
	Entities cedar.Entities `json:"entities,omitempty" validate:"omitempty"`
}

func (d AuthorizedDto) Validate() (fes []errs.FieldErr) {
	if d.Request.Principal.IsZero() {
		fes = append(fes, *errs.NewFieldErr("principal", "principal is required", "required"))
	}
	if d.Request.Action.IsZero() {
		fes = append(fes, *errs.NewFieldErr("action", "action is required", "required"))
	}
	if d.Request.Resource.IsZero() {
		fes = append(fes, *errs.NewFieldErr("resource", "resource is required", "required"))
	}

	return fes
}
