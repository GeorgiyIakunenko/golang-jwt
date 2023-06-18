package response

import "net/http"

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Basic struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendOk(w http.ResponseWriter, data any) {
	SendJson(w, http.StatusOK, data)
}
