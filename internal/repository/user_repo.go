package repository

import "github.com/ardrianenriquez/go-ecommerce/internal/domain"

type UserRepository interface {
	Create(u *domain.User) error
}
