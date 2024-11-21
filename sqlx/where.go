package sqlx

import (
	"fmt"
	"strings"
)

type _ConditionKind string

const (
	and _ConditionKind = "AND"
	or  _ConditionKind = "OR"
)

type _Condition struct {
	kind  _ConditionKind
	value Sqlizer
}

type WhereClause struct {
	depth      uint
	predicate  Sqlizer
	conditions []_Condition
}

func Where(predicate any) *WhereClause {
	return &WhereClause{0, &Sql{predicate}, []_Condition{}}
}

func (self *WhereClause) And(predicate any) *WhereClause {
	self.conditions = append(self.conditions, _Condition{
		kind:  and,
		value: &Sql{predicate},
	})

	return self
}

func (self *WhereClause) Or(predicate any) *WhereClause {
	self.conditions = append(self.conditions, _Condition{
		kind:  or,
		value: &Sql{predicate},
	})

	return self
}

func (self WhereClause) Sql() string {
	parts := []string{self.predicate.Sql()}

	for _, condition := range self.conditions {
		parts = append(parts, string(condition.kind)+" "+condition.value.Sql())
	}

	sql := strings.Join(parts, " ")

	if self.depth > 0 {
		sql = fmt.Sprintf("(%s)", sql)
	}

	return sql
}

func (self WhereClause) SqlPretty(indent string) string {
	parts := strings.Split(self.predicate.SqlPretty(indent), "\n")

	for _, condition := range self.conditions {
		parts = append(parts, string(condition.kind)+" "+condition.value.SqlPretty(indent))
	}

	if self.depth > 0 {
		for i := range parts {
			parts[i] = indent + parts[i]
		}

		parts = append([]string{"("}, parts...)
		parts = append(parts, ")")
	}

	return strings.Join(parts, "\n")
}

func (self *WhereClause) setDepth(depth uint) {
	self.depth = depth
	self.predicate.setDepth(depth + 1)

	for _, condition := range self.conditions {
		condition.value.setDepth(depth + 1)
	}
}
