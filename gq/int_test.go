package gq_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/gq"
)

func TestInt(t *testing.T) {
	t.Run("should resolve int", func(t *testing.T) {
		schema := gq.Int{}
		res := schema.Do(&gq.DoParams{
			Value: 69,
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(int)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != 69 {
			t.Fatalf("expected `%v`, received `%v`", 69, v)
		}
	})

	t.Run("should resolve int16", func(t *testing.T) {
		schema := gq.Int{}
		res := schema.Do(&gq.DoParams{
			Value: int16(69),
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(int16)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != 69 {
			t.Fatalf("expected `%v`, received `%v`", 69, v)
		}
	})

	t.Run("should resolve int32", func(t *testing.T) {
		schema := gq.Int{}
		res := schema.Do(&gq.DoParams{
			Value: int32(69),
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(int32)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != 69 {
			t.Fatalf("expected `%v`, received `%v`", 69, v)
		}
	})

	t.Run("should resolve int64", func(t *testing.T) {
		schema := gq.Int{}
		res := schema.Do(&gq.DoParams{
			Value: int64(69),
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(int64)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != 69 {
			t.Fatalf("expected `%v`, received `%v`", 69, v)
		}
	})

	t.Run("should not resolve", func(t *testing.T) {
		schema := gq.Int{}
		res := schema.Do(&gq.DoParams{
			Value: "test",
		})

		if res.Error == nil {
			t.FailNow()
		}
	})

	t.Run("should json", func(t *testing.T) {
		schema := gq.Int{}
		b, _ := json.Marshal(schema)

		if string(b) != `"int"` {
			t.FailNow()
		}
	})
}
