package handler

import (
	"codebase/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user usecase.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	errString := h.UserUsecase.CreateUser(user)
	w.Write([]byte(errString))
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user usecase.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	errString := h.UserUsecase.UpdateUser(user)
	w.Write([]byte(errString))
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users := h.UserUsecase.GetAllUser()
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	users := h.UserUsecase.GetUserID(userID)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user usecase.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	errString := h.UserUsecase.DeleteUser(user)
	w.Write([]byte(errString))
}
