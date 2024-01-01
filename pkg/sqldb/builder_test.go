package sqldb_test

import (
	"testing"

	"github.com/defryheryanto/nebula/pkg/sqldb"
	"github.com/stretchr/testify/assert"
)

type components struct {
	mainQuery string
	builder   *sqldb.QueryBuilder
}

func setupTest() *components {
	mainQuery := "SELECT * FROM users u"
	builder := sqldb.New(mainQuery)

	return &components{
		mainQuery: mainQuery,
		builder:   builder,
	}
}

func TestQueryBuilder_EmptyWhere(t *testing.T) {
	t.Parallel()

	s := setupTest()

	s.builder.Order("username DESC")

	query, args := s.builder.GetQuery(sqldb.AndOperator)
	assert.Equal(t, s.mainQuery+" ORDER BY username DESC", query)
	assert.Empty(t, args)
}

func TestQueryBuilder_WhereIn(t *testing.T) {
	t.Parallel()
	s := setupTest()

	parameters := []any{
		"hehe",
		"hoho",
	}
	s.builder.WhereIn("username IN ?", parameters)
	s.builder.Where("password = ?", "hihi")
	s.builder.Order("username DESC")

	query, args := s.builder.GetQuery(sqldb.AndOperator)
	assert.Equal(t, s.mainQuery+" WHERE username IN ($1,$2) AND password = $3 ORDER BY username DESC", query)
	assert.Equal(t, []any{
		"hehe",
		"hoho",
		"hihi",
	}, args)
}

func TestQueryBuilder_AndOperator(t *testing.T) {
	t.Parallel()
	s := setupTest()

	s.builder.Where("username = ?", "%hehe%")
	s.builder.Where("password = ?", "%huhu%")
	s.builder.Where("is_active = ?", true)
	s.builder.Order("username DESC")

	query, args := s.builder.GetQuery(sqldb.AndOperator)
	assert.Equal(t, s.mainQuery+" WHERE username = $1 AND password = $2 AND is_active = $3 ORDER BY username DESC", query)
	assert.Equal(t, []any{
		"%hehe%",
		"%huhu%",
		true,
	}, args)
}

func TestQueryBuilder_OrOperator(t *testing.T) {
	t.Parallel()
	s := setupTest()

	s.builder.Where("username = $1", "%hehe%")
	s.builder.Where("password = $2", "%huhu%")
	s.builder.Where("is_active = $3", true)
	s.builder.Order("username DESC")

	query, args := s.builder.GetQuery(sqldb.OrOperator)
	assert.Equal(t, s.mainQuery+" WHERE username = $1 OR password = $2 OR is_active = $3 ORDER BY username DESC", query)
	assert.Equal(t, []any{
		"%hehe%",
		"%huhu%",
		true,
	}, args)
}

func TestQueryBuilder_Join(t *testing.T) {
	t.Parallel()
	s := setupTest()

	parameters := []any{
		"hehe",
		"hoho",
	}
	s.builder.WhereIn("username IN ?", parameters)
	s.builder.Join("INNER JOIN user_groups ug ON u.group_id = ug.id")
	s.builder.Join("INNER JOIN user_group_permissions ugp ON ug.id = ugp.group_id")
	s.builder.Where("password = ?", "%hehe%")
	s.builder.Order("username DESC")

	query, _ := s.builder.GetQuery(sqldb.OrOperator)
	assert.Equal(
		t,
		s.mainQuery+
			" INNER JOIN user_groups ug ON u.group_id = ug.id "+
			"INNER JOIN user_group_permissions ugp ON ug.id = ugp.group_id "+
			"WHERE username IN ($1,$2) OR password = $3 ORDER BY username DESC",
		query,
	)
}

func TestQueryBuilder_LimitOffset(t *testing.T) {
	t.Parallel()
	s := setupTest()

	s.builder.Join("INNER JOIN user_groups ug ON u.group_id = ug.id")
	s.builder.Join("INNER JOIN user_group_permissions ugp ON ug.id = ugp.group_id")
	s.builder.Where("username = $1", "%hehe%")
	s.builder.Order("username DESC")
	s.builder.LimitOffset(10, 20)

	query, _ := s.builder.GetQuery(sqldb.OrOperator)
	assert.Equal(
		t,
		s.mainQuery+
			" INNER JOIN user_groups ug ON u.group_id = ug.id "+
			"INNER JOIN user_group_permissions ugp ON ug.id = ugp.group_id "+
			"WHERE username = $1 ORDER BY username DESC "+
			"LIMIT 10 OFFSET 20",
		query,
	)
}
