```go
func main() {
    var m sync.Map
    m.Store("key", 1)
    fmt.Println(m.Load("key")) // Output: 1

    go func() {
        m.Store("key", 2)
    }()

    time.Sleep(100 * time.Millisecond) // Allow goroutine to execute
    fmt.Println(m.Load("key")) // Output might be 1 or 2.  Race condition!
}
```