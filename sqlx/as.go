package sqlx

import "fmt"

type AsClause struct {
	depth uint
	stmt  Sqlizer
	alias string
}

func As(stmt Sqlizer, alias string) *AsClause {
	return &AsClause{0, stmt, alias}
}

func (self AsClause) Sql() string {
	return fmt.Sprintf("%s as \"%s\"", self.stmt.Sql(), self.alias)
}

func (self AsClause) SqlPretty() string {
	return fmt.Sprintf("%s as \"%s\"", self.stmt.SqlPretty(), self.alias)
}

func (self *AsClause) setDepth(depth uint) {
	self.depth = depth
	self.stmt.setDepth(depth + 1)
}
