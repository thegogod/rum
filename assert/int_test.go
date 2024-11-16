package assert_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/assert"
)

func TestInt(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := assert.Int().Required().Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := assert.Int().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("enum", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := assert.Int().Enum(1).Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := assert.Int().Enum(1).Validate(2)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := assert.Int().Required().Message("a test message").Validate(nil)

			if err == nil {
				t.FailNow()
			}

			if err.Error() != `{"errors":[{"rule":"required","message":"a test message"}]}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"errors":[{"rule":"required","message":"required"}]}`,
					err.Error(),
				)
			}
		})
	})

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := assert.Int().Min(5).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when gt min", func(t *testing.T) {
			err := assert.Int().Min(5).Validate(6)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when lt min", func(t *testing.T) {
			err := assert.Int().Min(5).Validate(4)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := assert.Int().Max(5).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when lt max", func(t *testing.T) {
			err := assert.Int().Max(5).Validate(4)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when gt max", func(t *testing.T) {
			err := assert.Int().Max(5).Validate(6)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := assert.Int().Min(1).Max(5)
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"int","min":1,"max":5}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"int","min":1,"max":5}`,
					string(b),
				)
			}
		})
	})
}

func BenchmarkInt(b *testing.B) {
	b.Run("int32", func(b *testing.B) {
		schema := assert.Int()

		for i := 0; i < b.N; i++ {
			err := schema.Validate(int32(1))

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("int64", func(b *testing.B) {
		schema := assert.Int()

		for i := 0; i < b.N; i++ {
			err := schema.Validate(int64(1))

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("min", func(b *testing.B) {
		schema := assert.Int().Min(5)

		for i := 0; i < b.N; i++ {
			err := schema.Validate(5)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("max", func(b *testing.B) {
		schema := assert.Int().Max(5)

		for i := 0; i < b.N; i++ {
			err := schema.Validate(5)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func ExampleInt() {
	schema := assert.Int()

	if err := schema.Validate(1); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(1); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}

func ExampleIntSchema_Min() {
	schema := assert.Int().Min(5)

	if err := schema.Validate(5); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(4); err != nil { // error
		panic(err)
	}
}

func ExampleIntSchema_Max() {
	schema := assert.Int().Max(5)

	if err := schema.Validate(5); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(6); err != nil { // error
		panic(err)
	}
}
