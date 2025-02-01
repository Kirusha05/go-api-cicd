package api

import (
	"testing"

	"github.com/Kirusha05/go-api-cicd/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	type testCase struct {
		expected []types.User
		err      error
	}

	usersService := NewUserService()

	t.Run("Get the list of users", func(t *testing.T) {
		tests := []testCase{
			{
				expected: []types.User{
					{Name: "Kirul", Email: "kiril@test.com", Age: 20},
				},
			},
		}

		for _, test := range tests {
			actual, err := usersService.GetUsers()
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		}
	})
}
