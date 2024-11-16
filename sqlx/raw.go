package sqlx

type RawStatement struct {
	stmt string
}

func Raw(stmt string) *RawStatement {
	return &RawStatement{stmt}
}

func (self RawStatement) Sql() string {
	return self.stmt
}

func (self RawStatement) SqlPretty() string {
	return self.stmt
}
