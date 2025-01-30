package api

import "github.com/Kirusha05/go-api-cicd/internal/types"

type UserService struct{}

func NewUserService() *UserService {
	userService := UserService{}
	return &userService
}

func (svc *UserService) GetUsers() ([]types.User, error) {
	users := []types.User{
		{
			Name:  "Kiril",
			Email: "kiril@test.com",
			Age:   20,
		},
	}
	return users, nil
}
