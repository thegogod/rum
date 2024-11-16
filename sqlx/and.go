package sqlx

import (
	"fmt"
	"strings"
)

type AndClause struct {
	depth      uint
	conditions []Sqlizer
}

func And(conditions ...Sqlizer) *AndClause {
	return &AndClause{0, conditions}
}

func (self AndClause) Sql() string {
	parts := []string{}

	for _, cond := range self.conditions {
		parts = append(parts, cond.Sql())
	}

	sql := strings.Join(parts, " AND ")

	if self.depth > 0 && len(parts) > 1 {
		sql = fmt.Sprintf("(%s)", sql)
	}

	return sql
}

func (self AndClause) SqlPretty() string {
	parts := []string{}

	for _, cond := range self.conditions {
		parts = append(parts, cond.Sql())
	}

	sql := strings.Join(parts, "\nAND ")

	if self.depth > 0 && len(parts) > 1 {
		sql = fmt.Sprintf("(%s)", sql)
	}

	return sql
}

func (self *AndClause) setDepth(depth uint) {
	self.depth = depth

	for _, cond := range self.conditions {
		cond.setDepth(self.depth + 1)
	}
}
