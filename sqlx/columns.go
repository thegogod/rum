package sqlx

import "strings"

type Columns []Sqlizer

func (self Columns) Sql() string {
	parts := []string{}

	for _, column := range self {
		parts = append(parts, column.Sql())
	}

	return strings.Join(parts, ", ")
}

func (self Columns) SqlPretty(indent string) string {
	parts := []string{}

	for i, column := range self {
		lines := strings.Split(column.SqlPretty(indent), "\n")

		for _, line := range lines {
			parts = append(parts, indent+line)
		}

		if i < len(self)-1 {
			parts[len(parts)-1] += ","
		}
	}

	return strings.Join(parts, "\n")
}

func (self Columns) setDepth(depth uint) {
	for _, column := range self {
		column.setDepth(depth)
	}
}
