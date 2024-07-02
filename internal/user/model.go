package user

import "time"

type User struct {
    ID             int64     `json:"id"`
    Username       string    `json:"username"`
    Email          string    `json:"email"`
    HashedPassword string    `json:"-"`
    Role           string    `json:"role"`
    CreatedAt      time.Time `json:"created_at"`
}
