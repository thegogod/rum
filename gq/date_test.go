package gq_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/thegogod/rum/gq"
)

func TestDate(t *testing.T) {
	t.Run("should resolve", func(t *testing.T) {
		schema := gq.Date{}
		res := schema.Do(&gq.DoParams{
			Value: time.Now(),
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		_, ok := res.Data.(time.Time)

		if !ok {
			t.Fatal(res.Data)
		}
	})

	t.Run("should not resolve", func(t *testing.T) {
		schema := gq.Date{}
		res := schema.Do(&gq.DoParams{
			Value: "test",
		})

		if res.Error == nil {
			t.FailNow()
		}
	})

	t.Run("should json", func(t *testing.T) {
		schema := gq.Date{}
		b, _ := json.Marshal(schema)

		if string(b) != `"Date"` {
			t.FailNow()
		}
	})
}
