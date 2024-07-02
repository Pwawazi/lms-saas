package user

import (
    "context"
    "database/sql"

    "github.com/Pwawazi/lms-saas/db/queries"
)

type Repository interface {
    CreateUser(ctx context.Context, user User) (User, error)
}

type repository struct {
    db *sql.DB
    queries *queries.Queries
}

func NewRepository(db *sql.DB) Repository {
    return &repository{
        db: db,
        queries: queries.New(db),
    }
}

func (r *repository) CreateUser(ctx context.Context, user User) (User, error) {
    createdUser, err := r.queries.CreateUser(ctx, queries.CreateUserParams{
        Username: user.Username,
        Email: user.Email,
        HashedPassword: user.HashedPassword,
        Role: user.Role,
    })
    if err != nil {
        return User{}, err
    }
    return User{
        ID: createdUser.ID,
        Username: createdUser.Username,
        Email: createdUser.Email,
        Role: createdUser.Role,
        CreatedAt: createdUser.CreatedAt,
    }, nil
}
