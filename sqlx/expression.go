package sqlx

import "strings"

type Expression struct {
	left  Sqlizer
	op    Sqlizer
	right Sqlizer
}

func Expr(left any, op any, right any) *Expression {
	return &Expression{&Sql{left}, &Sql{op}, &Sql{right}}
}

func (self Expression) Sql() string {
	parts := []string{}

	if self.left != nil {
		parts = append(parts, self.left.Sql())
	}

	if self.op != nil {
		parts = append(parts, self.op.Sql())
	}

	if self.right != nil {
		parts = append(parts, self.right.Sql())
	}

	return strings.Join(parts, " ")
}

func (self Expression) SqlPretty(indent string) string {
	parts := []string{}

	if self.left != nil {
		parts = append(parts, self.left.SqlPretty(indent))
	}

	if self.op != nil {
		parts = append(parts, self.op.SqlPretty(indent))
	}

	if self.right != nil {
		parts = append(parts, self.right.SqlPretty(indent))
	}

	return strings.Join(parts, " ")
}

func (self *Expression) setDepth(depth uint) {
	self.left.setDepth(depth)
	self.op.setDepth(depth)
	self.right.setDepth(depth)
}
