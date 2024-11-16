package sqlx

type RawStatement struct {
	depth uint
	stmt  string
}

func Raw(stmt string) *RawStatement {
	return &RawStatement{0, stmt}
}

func (self RawStatement) Sql() string {
	return self.stmt
}

func (self RawStatement) SqlPretty() string {
	return self.stmt
}

func (self *RawStatement) setDepth(depth uint) {
	self.depth = depth
}
