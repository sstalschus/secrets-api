package user_ctrl

import (
	"net/http"
)

type IController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}
