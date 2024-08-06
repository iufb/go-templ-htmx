package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/iufb/go-templ-htmx/service/auth"
	"github.com/iufb/go-templ-htmx/types"
	"github.com/iufb/go-templ-htmx/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.AuthUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	// compate pass
	if !auth.ValidatePassword(u.Password, payload.Password) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Wrong password."))
		return
	}
	utils.WriteJSON(w, http.StatusOK, types.Response{Message: "Logged successfully!"})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.AuthUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("User  with email %s already exists", payload.Email))
		return
	}
	// validate payload
	if err = utils.Validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", validationErrors))
		return
	}
	log.Println("Validate pass")
	hashedPass, err := auth.HashPass(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	// create user
	err = h.store.CreateUser(types.User{
		Email:    payload.Email,
		Password: hashedPass,
	})

	log.Println("Create pass")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, types.Response{Message: "Registered successfully"})
}
