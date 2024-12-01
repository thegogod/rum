package gq_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/gq"
)

func TestFloat(t *testing.T) {
	t.Run("should resolve float32", func(t *testing.T) {
		schema := gq.Float{}
		res := schema.Do(&gq.DoParams{
			Value: float32(1.22),
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(float32)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != 1.22 {
			t.Fatalf("expected `%v`, received `%v`", 1.22, v)
		}
	})

	t.Run("should resolve float64", func(t *testing.T) {
		schema := gq.Float{}
		res := schema.Do(&gq.DoParams{
			Value: 1.22,
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(float64)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != 1.22 {
			t.Fatalf("expected `%v`, received `%v`", 1.22, v)
		}
	})

	t.Run("should not resolve", func(t *testing.T) {
		schema := gq.Float{}
		res := schema.Do(&gq.DoParams{
			Value: "test",
		})

		if res.Error == nil {
			t.FailNow()
		}
	})

	t.Run("should json", func(t *testing.T) {
		schema := gq.Float{}
		b, _ := json.Marshal(schema)

		if string(b) != `"float"` {
			t.FailNow()
		}
	})
}
