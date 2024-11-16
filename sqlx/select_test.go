package sqlx_test

import (
	"testing"

	"github.com/thegogod/rum/sqlx"
)

func TestSelect(t *testing.T) {
	t.Run("from", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			sql := sqlx.Select("a", "b", "c").From("test").Sql()

			if sql != "SELECT a, b, c FROM test;" {
				t.Fatalf(sql)
			}
		})

		t.Run("select", func(t *testing.T) {
			sql := sqlx.Select("a", "b", "c").FromSelect(
				sqlx.Select("d", "e", "f").From("test"),
				"test",
			).Sql()

			if sql != "SELECT a, b, c FROM (SELECT d, e, f FROM test) as \"test\";" {
				t.Fatalf(sql)
			}
		})
	})

	t.Run("where", func(t *testing.T) {
		t.Run("and", func(t *testing.T) {
			sql := sqlx.Select(
				"a", "b", "c",
			).From("test").Where(
				sqlx.And(
					sqlx.Raw("a = b"),
					sqlx.Raw("b = c"),
				),
			).Sql()

			if sql != "SELECT a, b, c FROM test WHERE a = b AND b = c;" {
				t.Fatalf(sql)
			}
		})

		t.Run("or", func(t *testing.T) {
			sql := sqlx.Select(
				"a", "b", "c",
			).From("test").Where(
				sqlx.Or(
					sqlx.Raw("a = b"),
					sqlx.Raw("b = c"),
				),
			).Sql()

			if sql != "SELECT a, b, c FROM test WHERE a = b OR b = c;" {
				t.Fatalf(sql)
			}
		})

		t.Run("and or", func(t *testing.T) {
			sql := sqlx.Select(
				"a", "b", "c",
			).From("test").Where(
				sqlx.And(
					sqlx.And(
						sqlx.Raw("a = b"),
						sqlx.Raw("b = c"),
					),
					sqlx.Or(
						sqlx.Raw("a = b"),
						sqlx.Raw("b = c"),
					),
				),
			).Sql()

			if sql != "SELECT a, b, c FROM test WHERE (a = b AND b = c) AND (a = b OR b = c);" {
				t.Fatalf(sql)
			}
		})
	})
}
