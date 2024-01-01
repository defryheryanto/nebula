package repository

import (
	"context"
	"database/sql"

	"github.com/defryheryanto/nebula/internal/user"
	"github.com/defryheryanto/nebula/pkg/sqldb"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) First(ctx context.Context, filter *user.Filter) (*user.User, error) {
	builder := buildUserQuery(filter)

	builder.Order("users.id DESC")

	result := &User{}
	query, args := builder.GetQuery(sqldb.AndOperator)
	err := result.scan(r.db.QueryRowContext(ctx, query, args...))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result.convert(), nil
}
