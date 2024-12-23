package sqlx

type Sqlizer interface {
	Sql() string
	SqlPretty(indent string) string

	setDepth(depth uint)
}

type Sql struct {
	Value any
}

func Raw(value string) *Sql {
	return &Sql{value}
}

func (self Sql) Sql() string {
	switch v := self.Value.(type) {
	case string:
		return v
	case Sqlizer:
		return v.Sql()
	}

	panic("invalid type")
}

func (self Sql) SqlPretty(indent string) string {
	switch v := self.Value.(type) {
	case string:
		return v
	case Sqlizer:
		return v.SqlPretty(indent)
	}

	panic("invalid type")
}

func (self *Sql) setDepth(depth uint) {
	switch v := self.Value.(type) {
	case Sqlizer:
		v.setDepth(depth)
	}
}
