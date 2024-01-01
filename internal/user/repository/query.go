package repository

import (
	"github.com/defryheryanto/nebula/internal/user"
	"github.com/defryheryanto/nebula/pkg/sqldb"
)

const (
	selectUserQuery = `
		SELECT
			id,
			username,
			password,
			created_at,
			updated_at
		FROM public.users
	`
)

func buildUserQuery(filter *user.Filter) *sqldb.QueryBuilder {
	builder := sqldb.New(selectUserQuery)
	if filter.Username != "" {
		builder.Where("users.username = ?", filter.Username)
	}

	return builder
}
