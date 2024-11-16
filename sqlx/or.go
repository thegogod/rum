package sqlx

import (
	"fmt"
	"strings"
)

type OrClause struct {
	depth      uint
	conditions []Sqlizer
}

func Or(conditions ...Sqlizer) *OrClause {
	return &OrClause{0, conditions}
}

func (self OrClause) Sql() string {
	parts := []string{}

	for _, cond := range self.conditions {
		parts = append(parts, cond.Sql())
	}

	sql := strings.Join(parts, " OR ")

	if self.depth > 0 && len(parts) > 1 {
		sql = fmt.Sprintf("(%s)", sql)
	}

	return sql
}

func (self OrClause) SqlPretty() string {
	parts := []string{}

	for _, cond := range self.conditions {
		parts = append(parts, cond.Sql())
	}

	sql := strings.Join(parts, "\nOR ")

	if self.depth > 0 && len(parts) > 1 {
		sql = fmt.Sprintf("(%s)", sql)
	}

	return sql
}

func (self *OrClause) setDepth(depth uint) {
	self.depth = depth

	for _, cond := range self.conditions {
		cond.setDepth(self.depth + 1)
	}
}
