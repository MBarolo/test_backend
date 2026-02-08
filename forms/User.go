package forms

import "github.com/mbarolo/test_back/models"

type UserForm struct {
	Email          string `json:"email,omitempty"`
	HashedPassword string `json:"hashed_password,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
}

func (uf *UserForm) ToUser() *models.User {
	return &models.User{
		Email:          uf.Email,
		HashedPassword: uf.HashedPassword,
		FirstName:      uf.FirstName,
		LastName:       uf.LastName,
	}
}
