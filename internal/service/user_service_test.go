package service

import (
	"fmt"
	"testing"

	"github.com/ardrianenriquez/go-ecommerce/internal/domain"
)

type MockUserRepository struct {
	ErrToReturn error
}

func (m *MockUserRepository) Create(u *domain.User) error {
	return m.ErrToReturn
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name          string
		mockRepoError error
		wantErr       bool
	}{
		{
			name:          "Success",
			mockRepoError: nil,
			wantErr:       false,
		},
		{
			name:          "Failed",
			mockRepoError: fmt.Errorf("db connection failed"),
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := &MockUserRepository{
				ErrToReturn: tt.mockRepoError,
			}

			// Act
			service := NewUserService(mockRepo)
			user := &domain.User{
				Email: "sample@test.com",
			}

			err := service.RegisterUser(user)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expecting an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error occur: %v", err)
				}
			}
		})
	}
}
