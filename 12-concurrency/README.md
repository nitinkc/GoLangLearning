# Lesson 12: Concurrency

## What You'll Learn
- Goroutines
- Channels (unbuffered and buffered)
- Select statement
- WaitGroup and Mutex
- Common patterns

## Key Concepts

### Goroutines
```go
go myFunction()     // Start goroutine
go func() { }()     // Anonymous goroutine
```

### Channels
```go
ch := make(chan int)      // Unbuffered
ch := make(chan int, 10)  // Buffered

ch <- value   // Send
value := <-ch // Receive
close(ch)     // Close
```

### Select
```go
select {
case msg := <-ch1:
    // Handle ch1
case ch2 <- value:
    // Send to ch2
case <-time.After(time.Second):
    // Timeout
default:
    // Non-blocking
}
```

### Synchronization
```go
// WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()

// Mutex
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

## Running the Code
```bash
go run main.go
```
