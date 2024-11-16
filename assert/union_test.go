package assert_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/assert"
)

func TestUnion(t *testing.T) {
	t.Run("union", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := assert.Union(
				assert.String().Required(),
				assert.Int().Required(),
			).Validate("test")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := assert.Union(
				assert.String().Required(),
				assert.Int().Required(),
			).Validate(true)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := assert.Union(
				assert.String().Required(),
				assert.Int().Required(),
			).Message("a test message").Validate(true)

			if err == nil {
				t.FailNow()
			}

			if err.Error() != `{"errors":[{"rule":"type","message":"a test message"}]}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"errors":[{"rule":"type","message":"required"}]}`,
					err.Error(),
				)
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := assert.Union(
				assert.String().Required(),
				assert.Int().Required(),
			)

			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"union[string,int]"}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"union[string,int]"}`,
					string(b),
				)
			}
		})
	})
}

func BenchmarkUnion(b *testing.B) {
	b.Run("string or int", func(b *testing.B) {
		values := []any{"test", 1}
		schema := assert.Union(
			assert.String().Required(),
			assert.Int().Required(),
		)

		for i := 0; i < b.N; i++ {
			var err error

			if i%2 == 0 {
				err = schema.Validate(values[0])
			} else {
				err = schema.Validate(values[1])
			}

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func ExampleUnion() {
	schema := assert.Union(
		assert.String().Required(),
		assert.Int().Required(),
	)

	if err := schema.Validate("test"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(1); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(true); err != nil { // error
		panic(err)
	}
}
