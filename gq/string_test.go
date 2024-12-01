package gq_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/gq"
)

func TestString(t *testing.T) {
	t.Run("should resolve", func(t *testing.T) {
		schema := gq.String{}
		res := schema.Do(&gq.DoParams{
			Value: "testing123",
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(string)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != "testing123" {
			t.Fatalf("expected `%v`, received `%v`", "testing123", v)
		}
	})

	t.Run("should not resolve", func(t *testing.T) {
		schema := gq.String{}
		res := schema.Do(&gq.DoParams{
			Value: 123,
		})

		if res.Error == nil {
			t.FailNow()
		}
	})

	t.Run("should json", func(t *testing.T) {
		schema := gq.String{}
		b, _ := json.Marshal(schema)

		if string(b) != `"string"` {
			t.FailNow()
		}
	})
}
