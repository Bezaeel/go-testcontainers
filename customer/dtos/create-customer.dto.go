package dtos

type CustomerDTO struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}