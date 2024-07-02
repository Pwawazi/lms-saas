package user

import (
    "context"
    "errors"
    "golang.org/x/crypto/bcrypt"
)

type Service interface {
    CreateUser(ctx context.Context, username, email, password, role string) (User, error)
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo: repo}
}

func (s *service) CreateUser(ctx context.Context, username, email, password, role string) (User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return User{}, err
    }

    user := User{
        Username: username,
        Email: email,
        HashedPassword: string(hashedPassword),
        Role: role,
    }

    return s.repo.CreateUser(ctx, user)
}
