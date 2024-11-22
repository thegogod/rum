package sqlx

import (
	"fmt"
	"strings"
)

type JoinClause struct {
	method *string
	table  string
	where  *WhereClause
}

func Join(table string, predicate any) *JoinClause {
	return &JoinClause{nil, table, Where(predicate)}
}

func LeftJoin(table string, predicate any) *JoinClause {
	method := "LEFT"
	return &JoinClause{&method, table, Where(predicate)}
}

func LeftOuterJoin(table string, predicate any) *JoinClause {
	method := "LEFT OUTER"
	return &JoinClause{&method, table, Where(predicate)}
}

func RightJoin(table string, predicate any) *JoinClause {
	method := "RIGHT"
	return &JoinClause{&method, table, Where(predicate)}
}

func RightOuterJoin(table string, predicate any) *JoinClause {
	method := "RIGHT OUTER"
	return &JoinClause{&method, table, Where(predicate)}
}

func FullOuterJoin(table string, predicate any) *JoinClause {
	method := "FULL OUTER"
	return &JoinClause{&method, table, Where(predicate)}
}

func CrossJoin(table string, predicate any) *JoinClause {
	method := "CROSS"
	return &JoinClause{&method, table, Where(predicate)}
}

func (self *JoinClause) And(predicate any) *JoinClause {
	self.where.And(predicate)
	return self
}

func (self *JoinClause) Or(predicate any) *JoinClause {
	self.where.Or(predicate)
	return self
}

func (self JoinClause) Sql() string {
	parts := []string{}

	if self.method != nil {
		parts = append(parts, *self.method)
	}

	parts = append(parts, "JOIN", self.table, "ON")
	parts = append(parts, self.where.Sql())
	return strings.Join(parts, " ")
}

func (self JoinClause) SqlPretty(indent string) string {
	parts := []string{}

	if self.method != nil {
		parts = append(parts, fmt.Sprintf("%s JOIN %s", *self.method, self.table))
	} else {
		parts = append(parts, fmt.Sprintf("JOIN %s", self.table))
	}

	lines := strings.Split(self.where.SqlPretty(indent), "\n")
	parts = append(parts, indent+"ON "+lines[0])

	for _, line := range lines[1:] {
		parts = append(parts, indent+line)
	}

	return strings.Join(parts, "\n")
}

func (self *JoinClause) setDepth(_ uint) {

}
