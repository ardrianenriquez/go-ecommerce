package service

import (
	"github.com/ardrianenriquez/go-ecommerce/internal/domain"
	"github.com/ardrianenriquez/go-ecommerce/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) RegisterUser(u *domain.User) error {
	return s.repo.Create(u)
}
