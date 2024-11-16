# Box

a dependency injection library

# Usage

```go
b := box.New()

// put some things in the box
b.Put(&ServiceA{}, &ServiceB{})

// get something out of the box
box.Get[*ServiceA](b)
```

# Benchmarks

- coming soon!
