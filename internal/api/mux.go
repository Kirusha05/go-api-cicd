package api

import (
	"log"
	"net/http"

	"github.com/Kirusha05/go-api-cicd/internal/utils"

	"github.com/gorilla/mux"
)

type Mux struct {
	*mux.Router
	userService *UserService
}

func NewMux() *Mux {
	newMux := Mux{
		Router:      mux.NewRouter(),
		userService: NewUserService(),
	}

	newMux.HandleFunc("/", newMux.ListUsers).Methods("GET")

	return &newMux
}

func (m *Mux) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := m.userService.GetUsers()
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Could not get users")
	}

	log.Default().Println("users", users)
	utils.WriteJSONResponse(w, http.StatusOK, users)
}
