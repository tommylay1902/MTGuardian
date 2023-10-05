package dtos

type TodoDTO struct {
	Todo      *string `json:"todo"`
	Completed *bool   `json:"completed"`
}
