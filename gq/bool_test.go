package gq_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/gq"
)

func TestBool(t *testing.T) {
	t.Run("should resolve", func(t *testing.T) {
		schema := gq.Bool{}
		res := schema.Do(&gq.DoParams{
			Value: true,
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(bool)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != true {
			t.Fatalf("expected `%v`, received `%v`", true, v)
		}
	})

	t.Run("should not resolve", func(t *testing.T) {
		schema := gq.Bool{}
		res := schema.Do(&gq.DoParams{
			Value: "test",
		})

		if res.Error == nil {
			t.FailNow()
		}
	})

	t.Run("should json", func(t *testing.T) {
		schema := gq.Bool{}
		b, _ := json.Marshal(schema)

		if string(b) != `"bool"` {
			t.FailNow()
		}
	})
}
