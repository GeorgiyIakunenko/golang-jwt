package repo

import (
	"errors"
	"golang-jwt/model"
	"golang-jwt/server/request"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	RegisterUser(u request.RegisterRequest) (*model.User, error)
}

type UserRepository struct {
	users *[]*model.User
}

var (
	p1, _ = bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)
	p2, _ = bcrypt.GenerateFromPassword([]byte("22222222"), bcrypt.DefaultCost)

	users = &[]*model.User{
		&model.User{
			ID:       1,
			Email:    "alex@example.com",
			Name:     "Alex",
			Password: string(p1),
		},
		&model.User{
			ID:       2,
			Email:    "mary@example.com",
			Name:     "Mary",
			Password: string(p2),
		},
	}
)

func NewUserRepository() IUserRepository {
	return &UserRepository{users: users}
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	for _, user := range *r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	for _, user := range *r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) RegisterUser(u request.RegisterRequest) (*model.User, error) {

	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	for _, user := range *r.users {
		if user.Email == u.Email {
			return nil, errors.New("user is already exists")
		}
	}

	newUser := &model.User{
		ID:       len(*r.users) + 1,
		Email:    u.Email,
		Name:     u.Name,
		Password: string(password),
	}

	*r.users = append(*r.users, newUser)

	return newUser, nil
}
