package box_test

import (
	"testing"

	"github.com/thegogod/rum/box"
)

type A struct {
	Hello string
}

type B struct {
	World string
}

type C struct {
	Test string
}

func TestBox(t *testing.T) {
	t.Run("should inject", func(t *testing.T) {
		b := box.New()
		b.Put(&A{"world"}, &B{"hello"})

		fn, err := b.Inject(func(a *A, b *B) {
			if a.Hello != "world" || b.World != "hello" {
				t.Fatal(a, b)
			}
		})

		if err != nil {
			t.Fatal(err)
		}

		fn()
	})

	t.Run("should error when handler nil", func(t *testing.T) {
		b := box.New()
		_, err := b.Inject(nil)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("should error when handler not function", func(t *testing.T) {
		b := box.New()
		_, err := b.Inject("test")

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("should error when called with wrong argument types", func(t *testing.T) {
		b := box.New()
		b.Put(&A{"world"}, &A{"hello"})

		_, err := b.Inject(func(a *A, b *B) {
			if a.Hello != "world" || b.World != "hello" {
				t.Fatal(a, b)
			}
		})

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("should error when called with wrong number of arguments", func(t *testing.T) {
		b := box.New()
		b.Put(&A{"world"})

		_, err := b.Inject(func(a *A, b *B) {
			if a.Hello != "world" || b.World != "hello" {
				t.Fatal(a, b)
			}
		})

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("should get value", func(t *testing.T) {
		b := box.New()
		b.PutByKey("test", &A{"world"})

		if _, ok := b.Value("test").(*A); !ok {
			t.Fatal("expected singleton value")
		}
	})

	t.Run("should get value by type", func(t *testing.T) {
		b := box.New()
		b.Put(&A{"world"})

		a := box.Get[*A](b)

		if a == nil {
			t.Fatalf("expected singleton value")
		}
	})

	t.Run("should get value by key", func(t *testing.T) {
		b := box.New()
		b.PutByKey("test", &A{"world"})

		a := box.GetPath[*A](b, "test")

		if a == nil {
			t.Fatalf("expected singleton value")
		}
	})

	t.Run("should get value by path", func(t *testing.T) {
		b := box.New()
		a := box.New()
		a.PutByKey("hello", &A{Hello: "world"})
		b.PutByKey("test", a)

		c := box.GetPath[*A](b, "test", "hello")

		if c == nil {
			t.Fatalf("expected singleton value")
		}
	})

	t.Run("should not have deadline", func(t *testing.T) {
		b := box.New()

		if _, ok := b.Deadline(); ok {
			t.Fatalf("expected no deadline")
		}
	})

	t.Run("should not have error", func(t *testing.T) {
		b := box.New()

		if err := b.Err(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should fork", func(t *testing.T) {
		b := box.New()
		b.Put(&A{"world"}, &B{"hello"})

		fork := b.Fork()
		fork.Put(&C{"test"})

		if b.Len() != 2 {
			t.Fatalf("first box should have %d items, received %d", 2, b.Len())
		}

		if fork.Len() != 3 {
			t.Fatalf("second box should have %d items, received %d", 3, fork.Len())
		}
	})
}

func BenchmarkBox(b *testing.B) {
	container := box.New()
	container.Put(&A{"a"}, &B{"b"}, &C{"c"})
	fn, _ := container.Inject(func(a *A, b *B, c *C) {

	})

	for i := 0; i < b.N; i++ {
		fn()
	}
}
