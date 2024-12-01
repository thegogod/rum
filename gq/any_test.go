package gq_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/gq"
)

func TestAny(t *testing.T) {
	t.Run("should resolve", func(t *testing.T) {
		schema := gq.Any{}
		res := schema.Do(&gq.DoParams{
			Value: "testing",
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(string)

		if !ok {
			t.Fatal(res.Data)
		}

		if v != "testing" {
			t.Fatalf("expected `%s`, received `%s`", "testing", v)
		}
	})

	t.Run("should json", func(t *testing.T) {
		schema := gq.Any{}
		b, _ := json.Marshal(schema)

		if string(b) != `"any"` {
			t.FailNow()
		}
	})
}
