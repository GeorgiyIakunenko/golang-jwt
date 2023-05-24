package server

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthHandler struct {
	cfg *Config
}

func NewAuthHandler(cfg *Config) *AuthHandler {

	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		req := new(LogginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := NewUserRepository().GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := GenerateToken(user.ID, h.cfg.AccessTokenLifetimeMinutes, h.cfg.AccessTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshToken, err := GenerateToken(user.ID, h.cfg.RefreshTokenLifetimeMinutes, h.cfg.RefreshTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := LoginResponse{
			AccessToken:  accessString, // secret token is one direction
			RefreshToken: refreshToken,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "only post is allowed", http.StatusMethodNotAllowed)
	}
}

type UserHandler struct {
	cfg *Config
}

func NewUserHandler(cfg *Config) *UserHandler {
	return &UserHandler{
		cfg: cfg,
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		AuthHeader := r.Header.Get("Authorization")
		tokenString := GetTokenFromBearerString(AuthHeader)

		claims, err := ValidateToken(tokenString, h.cfg.AccessTokenSecret)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		user, err := NewUserRepository().GetUserById(claims.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		resp := UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		} // separate structure

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resp) // encode retrn error , we should check it

	default:
		http.Error(w, "only get method allowed", http.StatusMethodNotAllowed)

	}
}
