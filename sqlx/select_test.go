package sqlx_test

import (
	"os"
	"strings"
	"testing"

	"github.com/thegogod/rum/sqlx"
)

func TestSelect(t *testing.T) {
	t.Run("column", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/column_string.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("a", "b", "c").Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("select", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/column_select.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select().ColumnAs(
				sqlx.Select("a", "b", "c").From("test"),
				"results",
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("string and select", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/column_string_select.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("1", "2").ColumnAs(
				sqlx.Select("a", "b", "c").From("test"),
				"results",
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})
	})

	t.Run("from", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/from_string.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("a", "b", "c").From("test").Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("select", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/from_select.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("a", "b", "c").FromSelect(
				sqlx.Select("d", "e", "f").From("test"),
				"test",
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})
	})

	t.Run("where", func(t *testing.T) {
		t.Run("and", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/where_and.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"a", "b", "c",
			).From("test").Where(
				"a = b",
			).And(
				"b = c",
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("or", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/where_or.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"a", "b", "c",
			).From("test").Where(
				"a = b",
			).Or(
				"b = c",
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("and or", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/where_and_or.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"a", "b", "c",
			).From("test").Where(
				"c = c",
			).And(
				sqlx.Where("a = b").And("b = c"),
			).And(
				sqlx.Where("a = b").Or("b = c"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})
	})

	t.Run("pretty", func(t *testing.T) {
		t.Run("column", func(t *testing.T) {
			t.Run("string", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/column_string_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("a", "b", "c").SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("select", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/column_select_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select().ColumnAs(
					sqlx.Select("a", "b", "c").From("test"),
					"results",
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("string and select", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/column_string_select_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("1", "2").ColumnAs(
					sqlx.Select("a", "b", "c").From("test"),
					"results",
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})
		})

		t.Run("from", func(t *testing.T) {
			t.Run("string", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/from_string_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("a", "b", "c").From("test").SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("select", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/from_select_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("a", "b", "c").FromSelect(
					sqlx.Select("d", "e", "f").From("test"),
					"test",
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})
		})

		t.Run("where", func(t *testing.T) {
			t.Run("and", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/where_and_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select(
					"a", "b", "c",
				).From("test").Where(
					"a = b",
				).And(
					"b = c",
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("or", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/where_or_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select(
					"a", "b", "c",
				).From("test").Where(
					"a = b",
				).Or(
					"b = c",
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("and or", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/where_and_or_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select(
					"a", "b", "c",
				).From("test").Where(
					"c = c",
				).And(
					sqlx.Where("a = b").And("b = c"),
				).And(
					sqlx.Where("a = b").Or("b = c"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})
		})
	})
}
