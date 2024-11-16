package assert_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/assert"
)

func TestArray(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := assert.Array(assert.String()).Required().Validate([]any{})

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := assert.Array(assert.String()).Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("items", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := assert.Array(assert.String().Required()).Required().Validate([]string{"test"})

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when item not required", func(t *testing.T) {
			err := assert.Array(assert.String()).Required().Validate([]any{nil})

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when not array", func(t *testing.T) {
			err := assert.Array(assert.Bool()).Validate("test")

			if err == nil {
				t.FailNow()
			}
		})

		t.Run("should fail when item type invalid", func(t *testing.T) {
			err := assert.Array(assert.Bool()).Required().Validate([]string{"test"})

			if err == nil {
				t.FailNow()
			}
		})

		t.Run("should fail when item required", func(t *testing.T) {
			err := assert.Array(assert.String().Required()).Required().Validate([]any{nil})

			if err == nil {
				t.FailNow()
			}
		})
	})

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := assert.Array(assert.String()).Required().Message("a test message").Validate(nil)

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
			err := assert.Array(assert.String()).Min(5).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when gt min", func(t *testing.T) {
			err := assert.Array(assert.String()).Min(5).Validate([]string{
				"a", "b", "c", "d", "e",
			})

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when lt min", func(t *testing.T) {
			err := assert.Array(assert.String()).Min(5).Validate([]string{
				"a", "b", "c", "d",
			})

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := assert.Array(assert.String()).Max(5).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when lt max", func(t *testing.T) {
			err := assert.Array(assert.String()).Max(5).Validate([]string{
				"a", "b", "c", "d", "e",
			})

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when gt max", func(t *testing.T) {
			err := assert.Array(assert.String()).Max(5).Validate([]string{
				"a", "b", "c", "d", "e", "f",
			})

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := assert.Array(assert.String()).Required()
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"array[string]","items":{"type":"string"},"required":true}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"array[string]","items":{"type":"string"},"required":true}`,
					string(b),
				)
			}
		})
	})
}

func BenchmarkArray(b *testing.B) {
	b.Run("array", func(b *testing.B) {
		schema := assert.Array(assert.String())
		value := []string{"test"}

		for i := 0; i < b.N; i++ {
			err := schema.Validate(value)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("min", func(b *testing.B) {
		schema := assert.Array(assert.String()).Min(3)
		value := []string{"a", "b", "c"}

		for i := 0; i < b.N; i++ {
			err := schema.Validate(value)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("max", func(b *testing.B) {
		schema := assert.Array(assert.String()).Max(3)
		value := []string{"a", "b", "c"}

		for i := 0; i < b.N; i++ {
			err := schema.Validate(value)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func ExampleArray() {
	schema := assert.Array(assert.String().Required())

	if err := schema.Validate([]string{"test"}); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate([]int{1}); err != nil { // error
		panic(err)
	}
}
