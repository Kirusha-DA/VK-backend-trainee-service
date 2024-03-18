package dtos

type MovieDTO struct {
	ID          uint   `json:"-"`
	Name        string `json:"name"`
	Description string `json:"desctiption"`
	Rating      uint   `json:"rating"`
}
