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
			expected, err := os.ReadFile("./testcases/select/column/string.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select().
				Column("a").
				Column("b").
				Column("c").
				Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("select", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/column/select.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				sqlx.Select("a", "b", "c").From("test").As("results"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("string and select", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/column/string_select.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"1",
				"2",
				sqlx.Select("a", "b", "c").From("test").As("results"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})
	})

	t.Run("from", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/from/string.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("a", "b", "c").From("test").Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("select", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/from/select.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("a", "b", "c").From(
				sqlx.As(
					sqlx.Select("d", "e", "f").From("test"),
					"test",
				),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})
	})

	t.Run("where", func(t *testing.T) {
		t.Run("and", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/where/and.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"a", "b", "c",
			).From("test").Where(
				"a = b",
			).And(
				sqlx.Expr(
					sqlx.Select("*").From("tester"),
					"IS",
					"NULL",
				),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("or", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/where/or.sql")

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
			expected, err := os.ReadFile("./testcases/select/where/and_or.sql")

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

	t.Run("join", func(t *testing.T) {
		t.Run("join", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/join/join.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("*").From("a").Join(
				sqlx.Join("b", "a.id = b.id").And("b.deleted_at IS NULL"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("left", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/join/left_join.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("*").From("a").Join(
				sqlx.LeftJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("left outer", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/join/left_outer_join.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("*").From("a").Join(
				sqlx.LeftOuterJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("right", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/join/right_join.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("*").From("a").Join(
				sqlx.RightJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("right outer", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/join/right_outer_join.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("*").From("a").Join(
				sqlx.RightOuterJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("full outer", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/join/full_outer_join.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("*").From("a").Join(
				sqlx.FullOuterJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("cross", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/join/cross_join.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select("*").From("a").Join(
				sqlx.CrossJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
			).Sql()

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})
	})

	t.Run("group by", func(t *testing.T) {
		expected, err := os.ReadFile("./testcases/select/group_by.sql")

		if err != nil {
			t.Fatal(err)
		}

		sql := sqlx.Select(
			"id", "name", "created_at",
		).From("test").GroupBy("id").Sql()

		if sql != strings.TrimSuffix(string(expected), "\n") {
			t.Fatalf(sql)
		}
	})

	t.Run("limit offset", func(t *testing.T) {
		expected, err := os.ReadFile("./testcases/select/limit_offset.sql")

		if err != nil {
			t.Fatal(err)
		}

		sql := sqlx.Select(
			"*",
		).From("test").Limit("10").Offset("20").Sql()

		if sql != strings.TrimSuffix(string(expected), "\n") {
			t.Fatalf(sql)
		}
	})

	t.Run("order by", func(t *testing.T) {
		expected, err := os.ReadFile("./testcases/select/order_by.sql")

		if err != nil {
			t.Fatal(err)
		}

		sql := sqlx.Select(
			"*",
		).From("test").OrderBy(
			"created_at",
			sqlx.Desc,
		).Sql()

		if sql != strings.TrimSuffix(string(expected), "\n") {
			t.Fatalf(sql)
		}
	})

	t.Run("all", func(t *testing.T) {
		expected, err := os.ReadFile("./testcases/select/all.sql")

		if err != nil {
			t.Fatal(err)
		}

		sql := sqlx.Select(
			"a",
			"b",
			sqlx.Select("*").From("test").Limit("1").As("tester"),
		).From(
			"test",
		).Where(
			"a = 1",
		).Or(
			sqlx.Where("a = 2").And("b = 3"),
		).GroupBy("a").OrderBy("a", sqlx.Asc).Sql()

		if sql != strings.TrimSuffix(string(expected), "\n") {
			t.Fatalf(sql)
		}
	})

	t.Run("pretty", func(t *testing.T) {
		t.Run("column", func(t *testing.T) {
			t.Run("string", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/column/string_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select().
					Column("a").
					Column("b").
					Column("c").
					SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("select", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/column/select_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select(
					sqlx.Select("a", "b", "c").From("test").As("results"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("string and select", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/column/string_select_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select(
					"1",
					"2",
					sqlx.Select("a", "b", "c").From("test").As("results"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})
		})

		t.Run("from", func(t *testing.T) {
			t.Run("string", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/from/string_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("a", "b", "c").From("test").SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("select", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/from/select_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("a", "b", "c").From(
					sqlx.As(
						sqlx.Select("d", "e", "f").From("test"),
						"test",
					),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})
		})

		t.Run("where", func(t *testing.T) {
			t.Run("and", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/where/and_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select(
					"a", "b", "c",
				).From("test").Where(
					"a = b",
				).And(
					sqlx.Expr(
						sqlx.Select("*").From("tester"),
						"IS",
						"NULL",
					),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("or", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/where/or_pretty.sql")

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
				expected, err := os.ReadFile("./testcases/select/where/and_or_pretty.sql")

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

		t.Run("join", func(t *testing.T) {
			t.Run("join", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/join/join_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("*").From("a").Join(
					sqlx.Join("b", "a.id = b.id").And("b.deleted_at IS NULL"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("left", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/join/left_join_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("*").From("a").Join(
					sqlx.LeftJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("left outer", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/join/left_outer_join_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("*").From("a").Join(
					sqlx.LeftOuterJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("right", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/join/right_join_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("*").From("a").Join(
					sqlx.RightJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("right outer", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/join/right_outer_join_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("*").From("a").Join(
					sqlx.RightOuterJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("full outer", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/join/full_outer_join_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("*").From("a").Join(
					sqlx.FullOuterJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})

			t.Run("cross", func(t *testing.T) {
				expected, err := os.ReadFile("./testcases/select/join/cross_join_pretty.sql")

				if err != nil {
					t.Fatal(err)
				}

				sql := sqlx.Select("*").From("a").Join(
					sqlx.CrossJoin("b", "a.id = b.id").And("b.deleted_at IS NULL"),
				).SqlPretty("    ")

				if sql != strings.TrimSuffix(string(expected), "\n") {
					t.Fatalf(sql)
				}
			})
		})

		t.Run("group by", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/group_by_pretty.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"id", "name", "created_at",
			).From("test").GroupBy("id").SqlPretty("    ")

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("limit offset", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/limit_offset_pretty.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"*",
			).From("test").Limit("10").Offset("20").SqlPretty("    ")

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("order by", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/order_by_pretty.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"*",
			).From("test").OrderBy(
				"created_at",
				sqlx.Desc,
			).SqlPretty("    ")

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})

		t.Run("all", func(t *testing.T) {
			expected, err := os.ReadFile("./testcases/select/all_pretty.sql")

			if err != nil {
				t.Fatal(err)
			}

			sql := sqlx.Select(
				"a",
				"b",
				sqlx.Select("*").From("test").Limit("1").As("tester"),
			).From(
				"test",
			).Where(
				"a = 1",
			).Or(
				sqlx.Where("a = 2").And("b = 3"),
			).GroupBy("a").OrderBy("a", sqlx.Asc).SqlPretty("    ")

			if sql != strings.TrimSuffix(string(expected), "\n") {
				t.Fatalf(sql)
			}
		})
	})
}
