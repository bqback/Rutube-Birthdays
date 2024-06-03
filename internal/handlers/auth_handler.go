package handlers

import (
	"birthdays/internal/services"

	"net/http"
)

type AuthHandler struct {
	as services.IAuthService
	us services.IUserService
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (ah *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {

}

func (ah *AuthHandler) GetAuthService() services.IAuthService {
	return ah.as
}
