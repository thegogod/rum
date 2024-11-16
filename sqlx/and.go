package sqlx

import "strings"

type AndClause []Sqlizer

func And(conditions ...Sqlizer) AndClause {
	return conditions
}

func (self AndClause) Sql() string {
	parts := make([]string, len(self))

	for i, cond := range self {
		parts[i] = cond.Sql()
	}

	sql := strings.Join(parts, "")
	return sql
}
