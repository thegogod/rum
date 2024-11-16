# Box

a dependency injection library

# Usage

```go
b := box.New()
b.Put(&ServiceA{}, &ServiceB{})

fn, err := b.Inject(func (a *ServiceA, b *Service B) {
	fmt.Println(a, b)
})

if err != nil {
	panic(err)
}

fn()
```

# Benchmarks

- coming soon!
