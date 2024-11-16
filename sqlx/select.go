package sqlx

import (
	"fmt"
	"strings"
)

type SelectStatement struct {
	depth   uint
	columns []Sqlizer
	from    Sqlizer
	where   Sqlizer
}

func Select(columns ...string) *SelectStatement {
	cols := make([]Sqlizer, len(columns))

	for i, col := range columns {
		cols[i] = Raw(col)
	}

	return &SelectStatement{
		depth:   0,
		columns: cols,
		from:    nil,
		where:   nil,
	}
}

func (self *SelectStatement) Column(column string) *SelectStatement {
	self.columns = append(self.columns, Raw(column))
	return self
}

func (self *SelectStatement) ColumnAs(column string, alias string) *SelectStatement {
	self.columns = append(self.columns, As(Raw(column), alias))
	return self
}

func (self *SelectStatement) From(from string) *SelectStatement {
	self.from = Raw(from)
	return self
}

func (self *SelectStatement) FromSelect(stmt *SelectStatement, alias string) *SelectStatement {
	stmt.depth = self.depth + 1
	self.from = As(stmt, alias)
	return self
}

func (self *SelectStatement) Where(where Sqlizer) *SelectStatement {
	where.setDepth(self.depth)
	self.where = where
	return self
}

func (self SelectStatement) Sql() string {
	parts := []string{"SELECT"}
	columns := []string{}

	for _, column := range self.columns {
		columns = append(columns, column.Sql())
	}

	parts = append(parts, strings.Join(columns, ", "))

	if self.from != nil {
		parts = append(parts, "FROM", self.from.Sql())
	}

	if self.where != nil {
		parts = append(parts, "WHERE", self.where.Sql())
	}

	sql := strings.Join(parts, " ")

	if self.depth == 0 {
		sql += ";"
	} else {
		sql = fmt.Sprintf("(%s)", sql)
	}

	return sql
}

func (self SelectStatement) SqlPretty() string {
	return strings.Join([]string{}, "\n")
}

func (self *SelectStatement) setDepth(depth uint) {
	self.depth = depth
}
