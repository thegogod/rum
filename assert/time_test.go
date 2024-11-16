package assert_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/thegogod/rum/assert"
)

func TestTime(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := assert.Time().Required().Validate(time.Now().Format(time.RFC3339))

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

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := assert.Time().Required().Message("a test message").Validate(nil)

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

	t.Run("type", func(t *testing.T) {
		t.Run("should succeed when `time.Time`", func(t *testing.T) {
			err := assert.Time().Validate(time.Now())

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when time.Time string", func(t *testing.T) {
			err := assert.Time().Validate(time.Now().Format(time.RFC3339))

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when invalid string", func(t *testing.T) {
			err := assert.Time().Validate(time.Now().String())

			if err == nil {
				t.FailNow()
			}
		})

		t.Run("should fail when not string/time.Time", func(t *testing.T) {
			err := assert.Time().Validate(true)

			if err == nil {
				t.FailNow()
			}
		})
	})

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := assert.Time().Min(time.Now().AddDate(-1, 0, 0)).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when gt min", func(t *testing.T) {
			err := assert.Time().Min(time.Now().AddDate(-1, 0, 0)).Validate(time.Now().Format(time.RFC3339))

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when lt min", func(t *testing.T) {
			err := assert.Time().Min(time.Now()).Validate(time.Now().AddDate(-1, 0, 0).Format(time.RFC3339))

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := assert.Time().Max(time.Now().AddDate(1, 0, 0)).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when lt max", func(t *testing.T) {
			err := assert.Time().Max(time.Now().AddDate(1, 0, 0)).Validate(time.Now())

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when gt max", func(t *testing.T) {
			err := assert.Time().Max(time.Now()).Validate(time.Now().AddDate(1, 0, 0))

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := assert.Time()
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"time"}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"time"}`,
					string(b),
				)
			}
		})
	})
}

func BenchmarkTime(b *testing.B) {
	b.Run("time", func(b *testing.B) {
		schema := assert.Time()
		now := time.Now()

		for i := 0; i < b.N; i++ {
			err := schema.Validate(now)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("min", func(b *testing.B) {
		schema := assert.Time().Min(time.Now())
		now := time.Now()

		for i := 0; i < b.N; i++ {
			err := schema.Validate(now)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("max", func(b *testing.B) {
		schema := assert.Time().Max(time.Now())
		now := time.Now().Add(-1 * time.Hour)

		for i := 0; i < b.N; i++ {
			err := schema.Validate(now)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func ExampleTime() {
	schema := assert.Time()

	if err := schema.Validate(time.Now()); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(time.Now().Format(time.RFC3339)); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}

func ExampleTimeSchema_Min() {
	schema := assert.Time().Min(time.Now())

	if err := schema.Validate(time.Now().AddDate(-1, 0, 0)); err != nil { // error
		panic(err)
	}
}

func ExampleTimeSchema_Max() {
	schema := assert.Time().Max(time.Now())

	if err := schema.Validate(time.Now().AddDate(1, 0, 0)); err != nil { // error
		panic(err)
	}
}
