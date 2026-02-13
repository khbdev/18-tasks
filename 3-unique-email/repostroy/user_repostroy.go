package repostroy

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)



type User struct {
	ID int64
	Name string
	Email string
}

var ErrEmailAlreadyExists = errors.New("email a already exists")

type UserReposTory struct {
	db *sql.DB
}

func NewUserRepostry(db *sql.DB) *UserReposTory {
	return  &UserReposTory{db: db}
}

func (r *UserReposTory) Create(ctx context.Context, name, email string) (*User, error){
	email = strings.TrimSpace(strings.ToLower(email))

	const q = `
	INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id
	`
	var id int64
	err := r.db.QueryRowContext(ctx, 	q, name, email).Scan(&id)
	if err == nil {
		return  &User{ID: id, Name: name, Email: email}, nil
	}
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return  nil, ErrEmailAlreadyExists
	}
	return  nil, err
}

