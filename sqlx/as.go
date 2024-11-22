package sqlx

import (
	"fmt"
)

type AsClause struct {
	stmt  Sqlizer
	alias string
}

func As(stmt Sqlizer, alias string) *AsClause {
	return &AsClause{stmt, alias}
}

func (self AsClause) Sql() string {
	return fmt.Sprintf(`%s AS "%s"`, self.stmt.Sql(), self.alias)
}

func (self AsClause) SqlPretty(indent string) string {
	return fmt.Sprintf(`%s AS "%s"`, self.stmt.SqlPretty(indent), self.alias)
}

func (self *AsClause) setDepth(depth uint) {
	self.stmt.setDepth(depth)
}
