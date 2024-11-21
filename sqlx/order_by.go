package sqlx

type Direction string

const (
	Asc  Direction = "ASC"
	Desc Direction = "DESC"
)

type OrderByStatement struct {
	statement Sqlizer
	direction Direction
}

func OrderBy(statement any, direction Direction) *OrderByStatement {
	return &OrderByStatement{&Sql{statement}, direction}
}

func (self OrderByStatement) Sql() string {
	return self.statement.Sql() + " " + string(self.direction)
}

func (self OrderByStatement) SqlPretty(indent string) string {
	return self.statement.SqlPretty(indent) + " " + string(self.direction)
}

func (self *OrderByStatement) setDepth(depth uint) {
	self.statement.setDepth(depth)
}
