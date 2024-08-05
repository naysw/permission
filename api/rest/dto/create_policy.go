package dto

import (
	"strings"

	"github.com/naysw/permission/pkg/errs"
)

type CreatePolicy struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Document    string  `json:"document" validate:"required"`
}

func (d CreatePolicy) Validate() (fes []errs.FieldErr) {
	em := strings.TrimSpace(d.Name)
	doc := strings.TrimSpace(d.Document)
	switch {
	case em == "":
		fes = append(fes, *errs.NewFieldErr("name", "name is required", "required"))
	case len(doc) == 0:
		fes = append(fes, *errs.NewFieldErr("document", "document is required", "required"))
	}

	return fes
}
