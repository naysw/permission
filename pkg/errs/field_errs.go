package errs

import "fmt"

type FieldErr struct {
	Field    string      `json:"field"`
	Resource string      `json:"resource"`
	Code     string      `json:"code"`
	Value    interface{} `json:"value,omitempty"`
}

type FieldErrOption func(*FieldErr)

func WithValue(value interface{}) FieldErrOption {
	return func(e *FieldErr) {
		e.Value = value
	}
}

func NewFieldErr(
	field string,
	resource string,
	code string,
	opts ...FieldErrOption) *FieldErr {

	e := &FieldErr{
		Field:    field,
		Resource: resource,
		Code:     code,
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func NewMissing(field string) *FieldErr {
	return NewFieldErr(
		field,
		fmt.Sprintf("%s is missing", field),
		"missing",
	)
}
