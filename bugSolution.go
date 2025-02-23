```go
package main

import (


    "fmt"
    "sync"
    "time"
)

func main() {
    var m sync.Map
    var mu sync.Mutex
    m.Store("key", 1)
    fmt.Println(m.Load("key")) // Output: 1

    go func() {
        mu.Lock() 
        m.Store("key", 2)
        mu.Unlock()
    }()

    time.Sleep(100 * time.Millisecond) // Allow goroutine to execute
    fmt.Println(m.Load("key")) // Output: 2
}
```