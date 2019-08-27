monotonic - Let's Count, Concurrently
=====================================

```go
clock := monotonic.NewClock(0)

for _, thing := things {
  clock.Increment()
  go func() {
    thing.DoSomething()
    clock.Decrement()
  }()
}

for range time.NewTicker(5 * time.Second).C {
  n := c.Value()
  fmt.Printf("%d goroutines...\n", n)

  if n == 0 {
    break
  }
}
```
