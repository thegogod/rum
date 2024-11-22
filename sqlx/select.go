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
	joins   []Sqlizer
	groupBy Sqlizer
	orderBy Sqlizer
	limit   Sqlizer
	offset  Sqlizer
}

func Select(columns ...any) *SelectStatement {
	cols := make([]Sqlizer, len(columns))

	for i, col := range columns {
		cols[i] = &Sql{col}
	}

	return &SelectStatement{
		depth:   0,
		columns: cols,
		joins:   []Sqlizer{},
	}
}

func (self *SelectStatement) Column(column any) *SelectStatement {
	self.columns = append(self.columns, &Sql{column})
	return self
}

func (self *SelectStatement) From(from any) *SelectStatement {
	self.from = &Sql{from}
	return self
}

func (self *SelectStatement) Join(join Sqlizer) *SelectStatement {
	self.joins = append(self.joins, join)
	return self
}

func (self *SelectStatement) Where(predicate any) *SelectStatement {
	self.where = Where(predicate)
	return self
}

func (self *SelectStatement) And(predicates ...any) *SelectStatement {
	for _, predicate := range predicates {
		self.where.And(predicate)
	}

	return self
}

func (self *SelectStatement) Or(predicates ...any) *SelectStatement {
	for _, predicate := range predicates {
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

func (self *SelectStatement) As(alias string) Sqlizer {
	return As(self, alias)
}

func (self SelectStatement) Sql() string {
	self.setDepth(self.depth)
	parts := []string{"SELECT"}
	parts = append(parts, self.columns.Sql())

	if self.from != nil {
		parts = append(parts, "FROM", self.from.Sql())
	}

	for _, join := range self.joins {
		parts = append(parts, join.Sql())
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
	self.setDepth(self.depth)
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

	for _, join := range self.joins {
		lines := strings.Split(join.SqlPretty(indent), "\n")
		parts = append(parts, lines...)
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
