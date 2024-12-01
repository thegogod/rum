package gq_test

import (
	"encoding/json"
	"testing"

	"github.com/thegogod/rum/gq"
)

func TestRef(t *testing.T) {
	type User struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	ns := gq.New()
	ns.Register(gq.Object[User]{
		Name: "User",
		Fields: gq.Fields{
			"id":   gq.Field{},
			"name": gq.Field{},
			"email": gq.Field{
				Resolver: func(params *gq.ResolveParams) gq.Result {
					return gq.Result{Data: "test@gmail.com"}
				},
			},
		},
	})

	t.Run("should resolve", func(t *testing.T) {
		schema := ns.Ref("User")
		res := schema.Do(&gq.DoParams{
			Query: "{id,name,email}",
			Value: User{
				ID:   "1",
				Name: "test user",
			},
		})

		if res.Error != nil {
			t.Fatal(res.Error)
		}

		v, ok := res.Data.(User)

		if !ok {
			t.Fatal(res.Data)
		}

		if v.ID != "1" {
			t.Fatalf("expected `%v`, received `%v`", "1", v.ID)
		}

		if v.Name != "test user" {
			t.Fatalf("expected `%v`, received `%v`", "test user", v.ID)
		}

		if v.Email != "test@gmail.com" {
			t.Fatalf("expected `%v`, received `%v`", "test@gmail.com", v.ID)
		}
	})

	t.Run("should json", func(t *testing.T) {
		schema := ns.Ref("User")
		b, _ := json.Marshal(schema)

		if string(b) != `"User"` {
			t.FailNow()
		}
	})
}
