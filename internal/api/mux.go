package api

import (
	"encoding/json"
	"log"
	"my-api/internal/types"
	"net/http"

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
			Name:  "Kiril",
			Email: "kiril@test.com",
			Age:   20,
		},
	}
	log.Default().Println("users", users)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(users)
}
