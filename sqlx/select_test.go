package sqlx_test

import (
	"testing"

	"github.com/thegogod/rum/sqlx"
)

func TestSelect(t *testing.T) {
	t.Run("column", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			sql := sqlx.Select("a", "b", "c").Sql()

			if sql != "SELECT a, b, c;" {
				t.Fatalf(sql)
			}
		})

		t.Run("select", func(t *testing.T) {
			sql := sqlx.Select().ColumnSelect(
				sqlx.Select("a", "b", "c").From("test"),
				"results",
			).Sql()

			if sql != "SELECT (SELECT a, b, c FROM test) as \"results\";" {
				t.Fatalf(sql)
			}
		})

		t.Run("string and select", func(t *testing.T) {
			sql := sqlx.Select("1", "2").ColumnSelect(
				sqlx.Select("a", "b", "c").From("test"),
				"results",
			).Sql()

			if sql != "SELECT 1, 2, (SELECT a, b, c FROM test) as \"results\";" {
				t.Fatalf(sql)
			}
		})
	})

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
					sqlx.Raw("c = c"),
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

			if sql != "SELECT a, b, c FROM test WHERE c = c AND (a = b AND b = c) AND (a = b OR b = c);" {
				t.Fatalf(sql)
			}
		})
	})

	t.Run("pretty", func(t *testing.T) {
		t.Run("column", func(t *testing.T) {
			t.Run("string", func(t *testing.T) {
				sql := sqlx.Select("a", "b", "c").SqlPretty()

				if sql != "SELECT\n\ta,\n\tb,\n\tc;" {
					t.Fatalf(sql)
				}
			})

			t.Run("select", func(t *testing.T) {
				sql := sqlx.Select().ColumnSelect(
					sqlx.Select("a", "b", "c").From("test"),
					"results",
				).SqlPretty()

				if sql != "SELECT\n\t(SELECT a, b, c FROM test) as \"results\";" {
					t.Fatalf(sql)
				}
			})

			// t.Run("string and select", func(t *testing.T) {
			// 	sql := sqlx.Select("1", "2").ColumnSelect(
			// 		sqlx.Select("a", "b", "c").From("test"),
			// 		"results",
			// 	).Sql()

			// 	if sql != "SELECT 1, 2, (SELECT a, b, c FROM test) as \"results\";" {
			// 		t.Fatalf(sql)
			// 	}
			// })
		})
	})
}
