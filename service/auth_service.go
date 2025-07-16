package service

import (
	"test-naga-exchange/model"
	"test-naga-exchange/repository"
	"test-naga-exchange/util"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return &authService{r}
}

func (s *authService) Register(user *model.User) error {
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.ID = uuid.New()
	user.Password = hash
	user.Token = ""
	return s.repo.Create(user)
}

func (s *authService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || !util.CheckPassword(user.Password, password) {
		return nil, err
	}

	token, err := util.GenerateToken(email)
	if err != nil {
		return nil, err
	}
	user.Token = token
	s.repo.UpdateToken(user)
	return user, nil
}
