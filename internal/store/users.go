package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json: "email"`
	Passwrod  string `json: "-"`
	CreatedAt string `json: "created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
			INSERT INTO user (username, email, password)
			VALUES ($1, $2, $3) RETURNING id, created_at
		`

	if err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Passwrod).Scan(&user.ID, &user.CreatedAt); err != nil {
		return err
	}

	return nil
}
