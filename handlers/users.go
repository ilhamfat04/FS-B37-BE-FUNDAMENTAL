package handlers

import (
	dto "dumbmerch/dto/result"
	"dumbmerch/repositories"
	"encoding/json"
	"net/http"
)

// Declare handler struct here ...
type handler struct {
	UserRepository repositories.UserRepository
}

// Declare HandlerUser function here ...
func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

// Declare FindUsers method here ...
func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

// Declare GetUser method here ...

// Declare convertResponse function here ...
