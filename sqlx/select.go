package sqlx

import (
	"fmt"
	"strings"
)

type SelectStatement struct {
	depth   uint
	columns Columns
	from    Sqlizer
	where   *WhereClause
	groupBy Sqlizer
	orderBy Sqlizer
	limit   Sqlizer
	offset  Sqlizer
}

func Select(columns ...string) *SelectStatement {
	cols := make([]Sqlizer, len(columns))

	for i, col := range columns {
		cols[i] = Raw(col)
	}

	return &SelectStatement{
		depth:   0,
		columns: cols,
	}
}

func (self *SelectStatement) Column(column string) *SelectStatement {
	self.columns = append(self.columns, Raw(column))
	return self
}

func (self *SelectStatement) ColumnAs(column any, alias string) *SelectStatement {
	switch v := column.(type) {
	case string:
		self.columns = append(self.columns, As(&Sql{v}, alias))
		break
	case *SelectStatement:
		v.depth = self.depth + 1
		self.columns = append(self.columns, As(&Sql{v}, alias))
		break
	}

	return self
}

func (self *SelectStatement) From(from any) *SelectStatement {
	switch v := from.(type) {
	case string:
		self.from = &Sql{from}
		break
	case *SelectStatement:
		v.depth = self.depth + 1
		self.from = v
		break
	case *AsClause:
		switch v := v.stmt.(type) {
		case *SelectStatement:
			v.depth = self.depth + 1
			break
		}

		self.from = v
		break
	}

	return self
}

func (self *SelectStatement) Where(predicate any) *SelectStatement {
	switch v := predicate.(type) {
	case Sqlizer:
		v.setDepth(self.depth + 1)
	}

	self.where = Where(predicate)
	return self
}

func (self *SelectStatement) And(predicates ...any) *SelectStatement {
	for _, predicate := range predicates {
		switch v := predicate.(type) {
		case Sqlizer:
			v.setDepth(self.depth + 1)
		}

		self.where.And(predicate)
	}

	return self
}

func (self *SelectStatement) Or(predicates ...any) *SelectStatement {
	for _, predicate := range predicates {
		switch v := predicate.(type) {
		case Sqlizer:
			v.setDepth(self.depth + 1)
		}

		self.where.Or(predicate)
	}

	return self
}

func (self *SelectStatement) GroupBy(groupBy string) *SelectStatement {
	self.groupBy = &Sql{groupBy}
	return self
}

func (self *SelectStatement) OrderBy(statement any, direction Direction) *SelectStatement {
	self.orderBy = OrderBy(statement, direction)
	return self
}

func (self *SelectStatement) Limit(limit string) *SelectStatement {
	self.limit = &Sql{limit}
	return self
}

func (self *SelectStatement) Offset(offset string) *SelectStatement {
	self.offset = &Sql{offset}
	return self
}

func (self SelectStatement) Sql() string {
	parts := []string{"SELECT"}
	parts = append(parts, self.columns.Sql())

	if self.from != nil {
		parts = append(parts, "FROM", self.from.Sql())
	}

	if self.where != nil {
		parts = append(parts, "WHERE", self.where.Sql())
	}

	if self.groupBy != nil {
		parts = append(parts, "GROUP BY", self.groupBy.Sql())
	}

	if self.orderBy != nil {
		parts = append(parts, "ORDER BY", self.orderBy.Sql())
	}

	if self.limit != nil {
		parts = append(parts, "LIMIT", self.limit.Sql())
	}

	if self.offset != nil {
		parts = append(parts, "OFFSET", self.offset.Sql())
	}

	sql := strings.Join(parts, " ")

	if self.depth == 0 {
		sql += ";"
	} else {
		sql = fmt.Sprintf("(%s)", sql)
	}

	return sql
}

func (self SelectStatement) SqlPretty(indent string) string {
	parts := []string{}

	if self.depth > 0 {
		parts = append(parts, "(")
	}

	parts = append(parts, "SELECT")
	parts = append(
		parts,
		strings.Split(self.columns.SqlPretty(indent), "\n")...,
	)

	if self.from != nil {
		lines := strings.Split(self.from.SqlPretty(indent), "\n")
		parts = append(parts, "FROM "+lines[0])
		parts = append(parts, lines[1:]...)
	}

	if self.where != nil {
		lines := strings.Split(self.where.SqlPretty(indent), "\n")
		parts = append(parts, "WHERE "+lines[0])
		parts = append(parts, lines[1:]...)
	}

	if self.groupBy != nil {
		lines := strings.Split(self.groupBy.SqlPretty(indent), "\n")
		parts = append(parts, "GROUP BY "+lines[0])
		parts = append(parts, lines[1:]...)
	}

	if self.orderBy != nil {
		lines := strings.Split(self.orderBy.SqlPretty(indent), "\n")
		parts = append(parts, "ORDER BY "+lines[0])
		parts = append(parts, lines[1:]...)
	}

	if self.limit != nil {
		lines := strings.Split(self.limit.SqlPretty(indent), "\n")
		parts = append(parts, "LIMIT "+lines[0])
		parts = append(parts, lines[1:]...)
	}

	if self.offset != nil {
		lines := strings.Split(self.offset.SqlPretty(indent), "\n")
		parts = append(parts, "OFFSET "+lines[0])
		parts = append(parts, lines[1:]...)
	}

	if self.depth > 0 {
		for i := 1; i < len(parts); i++ {
			parts[i] = indent + parts[i]
		}

		parts = append(parts, ")")
	}

	sql := strings.Join(parts, "\n")

	if self.depth == 0 {
		sql += ";"
	}

	return sql
}

func (self *SelectStatement) setDepth(depth uint) {
	self.depth = depth
	self.columns.setDepth(depth + 1)

	if self.from != nil {
		self.from.setDepth(depth + 1)
	}

	if self.where != nil {
		self.where.setDepth(depth)
	}
}
