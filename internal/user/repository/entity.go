package repository

import (
	"time"

	"github.com/defryheryanto/nebula/internal/user"
	"github.com/defryheryanto/nebula/pkg/sqldb"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) scan(scanner sqldb.Scanner) error {
	return scanner.Scan(
		&u.ID,
		&u.Username,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (u *User) convert() *user.User {
	if u == nil {
		return nil
	}

	return &user.User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
