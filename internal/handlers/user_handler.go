package handlers

import (
	"birthdays/internal/services"

	"net/http"
)

type UserHandler struct {
	us services.IUserService
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (uh *UserHandler) Subscribe(w http.ResponseWriter, r *http.Request) {

}

func (uh *UserHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {

}
