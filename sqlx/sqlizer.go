package sqlx

type Sqlizer interface {
	Sql() string
	SqlPretty() string

	setDepth(depth uint)
}
