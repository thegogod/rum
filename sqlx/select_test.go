package sqlx_test

import (
	"testing"

	"github.com/thegogod/rum/sqlx"
)

func TestSelect(t *testing.T) {
	t.Run("should build basic", func(t *testing.T) {
		sql := sqlx.Select("a", "b", "c").From("test").Sql()

		if sql != "SELECT a, b, c FROM test;" {
			t.Fatalf(sql)
		}
	})

	t.Run("should build with nested from", func(t *testing.T) {
		sql := sqlx.Select("a", "b", "c").FromSelect(
			sqlx.Select("d", "e", "f").From("test"),
			"test",
		).Sql()

		if sql != "SELECT a, b, c FROM (SELECT d, e, f FROM test) as \"test\";" {
			t.Fatalf(sql)
		}
	})
}
