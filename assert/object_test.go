package assert_test

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/thegogod/rum/assert"
)

func TestObject(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := assert.Object().Field(
				"hello", assert.String().Enum("world").Required(),
			).Required().Validate(map[string]any{
				"hello": "world",
			})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should succeed when field nil and not required", func(t *testing.T) {
			err := assert.Object().Field(
				"hello", assert.String().Enum("world"),
			).Required().Validate(map[string]any{})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should fail when nil", func(t *testing.T) {
			err := assert.Object().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})

		t.Run("should fail when required field nil", func(t *testing.T) {
			err := assert.Object().Fields(map[string]assert.Schema{
				"hello": assert.String().Required(),
			}).Required().Validate(map[string]any{})

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := assert.Object().Required().Message("a test message").Validate(nil)

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

	t.Run("map", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := assert.Object().Field(
				"hello",
				assert.Object().Field("hello", assert.String().Enum("world")).Required(),
			).Required().Validate(map[string]any{
				"hello": map[string]any{
					"hello": "world",
				},
			})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should fail when schema not found", func(t *testing.T) {
			err := assert.Object().Field(
				"hello",
				assert.Object().Field("hello", assert.String().Enum("world").Required()).Required(),
			).Required().Validate(map[string]any{
				"hello": map[string]any{
					"hello1": "world",
				},
			})

			if err == nil {
				t.Fatal()
			}
		})

		t.Run("should fail when invalid schema", func(t *testing.T) {
			err := assert.Object().Field(
				"hello",
				assert.Object().Field("hello", assert.String().Enum("world")).Required(),
			).Required().Validate(map[string]any{
				"hello": "world",
			})

			if err == nil {
				t.Fatal()
			}
		})

		t.Run("should fail when missing required field", func(t *testing.T) {
			err := assert.Object().Field(
				"hello",
				assert.String().Required(),
			).Field(
				"world",
				assert.Bool().Required(),
			).Validate(map[string]any{
				"hello": "test",
			})

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("struct", func(t *testing.T) {
		type HelloWorld struct {
			Hello string `json:"hello"`
		}

		type Hello struct {
			Hello HelloWorld `json:"hello"`
		}

		t.Run("should succeed", func(t *testing.T) {
			err := assert.Object().Field(
				"hello",
				assert.Object().Field("hello", assert.String().Enum("world")).Required(),
			).Required().Validate(Hello{
				Hello: HelloWorld{"world"},
			})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should fail when invalid schema", func(t *testing.T) {
			err := assert.Object().Field(
				"hello",
				assert.Object().Field("hello", assert.String().Enum("world")).Required(),
			).Required().Validate(HelloWorld{
				Hello: "world",
			})

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("extend", func(t *testing.T) {
		t.Run("should add fields", func(t *testing.T) {
			a := assert.Object().Fields(map[string]assert.Schema{
				"a": assert.Int().Required(),
				"b": assert.Int().Required(),
				"c": assert.Int().Required(),
			})

			b := assert.Object().Fields(map[string]assert.Schema{
				"d": assert.String().Required(),
				"e": assert.Bool().Required(),
			})

			c := a.Extend(b)
			err := a.Validate(map[string]any{
				"a": 1,
				"b": 2,
				"c": 3,
			})

			if err != nil {
				t.Error(err)
			}

			err = b.Validate(map[string]any{
				"d": "test",
				"e": true,
			})

			if err != nil {
				t.Error(err)
			}

			err = c.Validate(map[string]any{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": "test",
				"e": true,
			})

			if err != nil {
				t.Error(err)
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := assert.Object().Field(
				"username", assert.String().Regex(regexp.MustCompile("^[0-9a-zA-Z_-]+$")).Required(),
			).Field(
				"password", assert.String().Min(5).Max(20).Required(),
			).Field(
				"staySignedIn", assert.Bool(),
			)

			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"object","fields":{"username":{"type":"string","regex":"^[0-9a-zA-Z_-]+$","required":true},"password":{"type":"string","min":5,"max":20,"required":true},"staySignedIn":{"type":"bool"}}}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"object","fields":{"username":{"type":"string","regex":"^[0-9a-zA-Z_-]+$","required":true},"password":{"type":"string","min":5,"max":20,"required":true},"staySignedIn":{"type":"bool"}}}`,
					string(b),
				)
			}
		})
	})
}

func BenchmarkObject(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		schema := assert.Object().Fields(map[string]assert.Schema{
			"email":    assert.String(),
			"password": assert.String(),
		})

		data := map[string]string{
			"email":    "test",
			"password": "test",
		}

		for i := 0; i < b.N; i++ {
			err := schema.Validate(data)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("array", func(b *testing.B) {
		schema := assert.Object().Fields(map[string]assert.Schema{
			"email":    assert.String(),
			"password": assert.String(),
			"phones":   assert.Array(assert.String()),
		})

		data := map[string]any{
			"email":    "test",
			"password": "test",
			"phonees":  []string{"1112223333"},
		}

		for i := 0; i < b.N; i++ {
			err := schema.Validate(data)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("object", func(b *testing.B) {
		schema := assert.Object().Fields(map[string]assert.Schema{
			"email":    assert.String(),
			"password": assert.String(),
			"created_by": assert.Object().Fields(map[string]assert.Schema{
				"email":    assert.String(),
				"password": assert.String(),
			}),
		})

		data := map[string]any{
			"email":    "test",
			"password": "test",
			"phonees": map[string]string{
				"email":    "test",
				"password": "test",
			},
		}

		for i := 0; i < b.N; i++ {
			err := schema.Validate(data)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func ExampleObject() {
	schema := assert.Object().Field(
		"email", assert.String().Email().Required(),
	).Field(
		"password", assert.String().Min(5).Max(20).Required(),
	)

	if err := schema.Validate(map[string]any{
		"email":    "test@test.com",
		"password": "mytestpassword",
	}); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(map[string]any{
		"email": "test@test.com",
	}); err != nil { // error
		panic(err)
	}
}
