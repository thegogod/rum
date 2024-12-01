package gq_test

import (
	"testing"
	"time"

	"github.com/thegogod/rum/gq"
)

func TestPointer(t *testing.T) {
	t.Run("should resolve string as pointer", func(t *testing.T) {
		schema := gq.Pointer{gq.String{}}
		res := schema.Do(&gq.DoParams{
			Value: "testing",
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(*string)

		if !ok {
			t.Fatal(res.Data)
		}

		if *v != "testing" {
			t.Fatalf("expected `%s`, received `%s`", "testing", *v)
		}
	})

	t.Run("should resolve bool as pointer", func(t *testing.T) {
		schema := gq.Pointer{gq.Bool{}}
		res := schema.Do(&gq.DoParams{
			Value: true,
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(*bool)

		if !ok {
			t.Fatal(res.Data)
		}

		if *v != true {
			t.Fatalf("expected `%v`, received `%v`", true, *v)
		}
	})

	t.Run("should resolve int as pointer", func(t *testing.T) {
		schema := gq.Pointer{gq.Int{}}
		res := schema.Do(&gq.DoParams{
			Value: 102,
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(*int)

		if !ok {
			t.Fatal(res.Data)
		}

		if *v != 102 {
			t.Fatalf("expected `%v`, received `%v`", 102, *v)
		}
	})

	t.Run("should resolve float as pointer", func(t *testing.T) {
		schema := gq.Pointer{gq.Float{}}
		res := schema.Do(&gq.DoParams{
			Value: 11.123,
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(*float64)

		if !ok {
			t.Fatal(res.Data)
		}

		if *v != 11.123 {
			t.Fatalf("expected `%v`, received `%v`", 11.123, *v)
		}
	})

	t.Run("should resolve `time.Time` as pointer", func(t *testing.T) {
		schema := gq.Pointer{gq.Date{}}
		res := schema.Do(&gq.DoParams{
			Value: time.Now(),
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		_, ok := res.Data.(*time.Time)

		if !ok {
			t.Fatal(res.Data)
		}
	})

	t.Run("should resolve struct as pointer", func(t *testing.T) {
		type User struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		schema := gq.Pointer{gq.Object[User]{
			Name: "User",
			Fields: gq.Fields{
				"name":  gq.Field{Type: gq.String{}},
				"email": gq.Field{Type: gq.String{}},
			},
		}}

		res := schema.Do(&gq.DoParams{
			Query: "{name,email}",
			Value: User{
				Name:  "test",
				Email: "test@test.com",
			},
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		user, ok := res.Data.(*User)

		if !ok {
			t.Fatal(res.Data)
		}

		if user.Name != "test" {
			t.Fatalf("expected `%s`, received `%s`", "test", user.Name)
		}

		if user.Email != "test@test.com" {
			t.Fatalf("expected `%s`, received `%s`", "test@test.com", user.Email)
		}
	})
}
