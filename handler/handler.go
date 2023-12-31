package handler

import (
	"encoding/json"
	"golang-jwt/config"
	"golang-jwt/repo"
	"golang-jwt/server/request"
	"golang-jwt/server/response"
	"golang-jwt/server/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(request.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.SendBadRequestError(w, err)
			return
		}

		user, err := repo.NewUserRepository().GetUserByEmail(req.Email)
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		accessString, err := service.GenerateToken(user.ID, h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
		if err != nil {
			response.SendServerError(w, err)
			return
		}

		refreshString, err := service.GenerateToken(user.ID, h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
		if err != nil {
			response.SendServerError(w, err)
			return
		}

		resp := response.TokenResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		response.SendOk(w, resp)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

type UserHandler struct {
	cfg *config.Config
}

type RefreshHandler struct {
	cfg *config.Config
}

func NewRefreshHandler(cfg *config.Config) *RefreshHandler {
	return &RefreshHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(request.RegisterRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.SendBadRequestError(w, err)
			return
		}

		_, err := repo.NewUserRepository().RegisterUser(*req)
		if err != nil {
			response.SendBadRequestError(w, err)
			return
		}

		response.SendOk(w, `user was successfully registered`)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *RefreshHandler) GetTokenPair(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		UserRefreshToken := r.Header.Get("Authorization")

		user, err := service.ValidateToken(UserRefreshToken, h.cfg.RefreshSecret)
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		accessToken, err := service.GenerateToken(user.ID, h.cfg.AccessLifetimeMinutes, h.cfg.AccessSecret)
		if err != nil {
			response.SendServerError(w, err)
			return
		}

		refreshToken, err := service.GenerateToken(user.ID, h.cfg.RefreshLifetimeMinutes, h.cfg.RefreshSecret)
		if err != nil {
			response.SendServerError(w, err)
			return
		}

		refreshResponse := response.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		response.SendOk(w, refreshResponse)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)

	}

}

func NewUserHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		claims, err := service.ValidateToken(service.GetTokenFromBearerString(r.Header.Get("Authorization")), h.cfg.AccessSecret)
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		user, err := repo.NewUserRepository().GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		resp := response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		response.SendOk(w, resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
