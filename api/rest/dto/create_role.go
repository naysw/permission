package dto

type CreateRole struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

func (c CreateRole) Validate() error {
	return nil
}
