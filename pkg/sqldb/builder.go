package sqldb

import (
	"fmt"
	"strings"
)

type Operator string

const (
	AndOperator Operator = "AND"
	OrOperator  Operator = "OR"
)

type QueryBuilder struct {
	mainQuery        string
	joinQuery        []string
	whereQuery       []string
	whereArgs        []any
	orderQuery       string
	limitOffsetQuery string
}

func New(mainQuery string) *QueryBuilder {
	return &QueryBuilder{
		mainQuery:        mainQuery,
		joinQuery:        []string{},
		whereQuery:       []string{},
		whereArgs:        []any{},
		orderQuery:       "",
		limitOffsetQuery: "",
	}
}

func (b *QueryBuilder) Where(query string, args any) {
	b.whereQuery = append(b.whereQuery, query)
	b.whereArgs = append(b.whereArgs, args)
}

func (b *QueryBuilder) WhereIn(query string, args []any) {
	inParametersArgs := make([]string, 0, len(args))
	for range args {
		inParametersArgs = append(inParametersArgs, "?")
	}

	inStatement := fmt.Sprintf("(%s)", strings.Join(inParametersArgs, ","))
	query = strings.ReplaceAll(query, "?", inStatement)

	b.whereQuery = append(b.whereQuery, query)
	b.whereArgs = append(b.whereArgs, args...)
}

func (b *QueryBuilder) Join(joinQuery string) {
	b.joinQuery = append(b.joinQuery, joinQuery)
}

func (b *QueryBuilder) Order(orderBy string) {
	b.orderQuery = fmt.Sprintf(" ORDER BY %s", orderBy)
}

func (b *QueryBuilder) LimitOffset(limit, offset int64) {
	b.limitOffsetQuery = fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
}

func (b *QueryBuilder) GetQuery(op Operator) (string, []any) {
	whereQuery := ""
	if len(b.whereQuery) > 0 {
		argIndex := 1
		for i, q := range b.whereQuery {
			parts := strings.Split(q, ",")
			if len(parts) > 1 {
				for range parts {
					b.whereQuery[i] = strings.Replace(b.whereQuery[i], "?", fmt.Sprintf("$%d", argIndex), 1)
					argIndex++
				}
				continue
			}

			b.whereQuery[i] = strings.ReplaceAll(q, "?", fmt.Sprintf("$%d", argIndex))
			argIndex++
		}
		whereQuery = " WHERE " + strings.Join(b.whereQuery, fmt.Sprintf(" %s ", string(op)))
	}

	joinQuery := ""
	if len(b.joinQuery) > 0 {
		joinQuery = " " + strings.Join(b.joinQuery, " ")
	}

	query := b.mainQuery + joinQuery + whereQuery + b.orderQuery + b.limitOffsetQuery

	return query, b.whereArgs
}
