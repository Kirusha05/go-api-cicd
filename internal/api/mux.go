package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kirusha05/go-api-cicd/internal/types"

	"github.com/gorilla/mux"
)

type Mux struct {
	*mux.Router
}

func NewMux() *Mux {
	newMux := Mux{
		Router: mux.NewRouter(),
	}

	newMux.HandleFunc("/", newMux.ListUsers).Methods("GET")

	return &newMux
}

func (m *Mux) ListUsers(w http.ResponseWriter, r *http.Request) {
	users := []types.User{
		{
			Name:  "Kirill",
			Email: "kirill@test.com",
			Age:   19,
		},
	}
	log.Default().Println("users", users)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(users)
}
