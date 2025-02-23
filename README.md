# Go sync.Map Race Condition Example

This repository demonstrates a potential race condition when using Go's `sync.Map` without proper synchronization mechanisms in concurrent operations.  The example shows how this can lead to unpredictable results.

## Bug

The `bug.go` file contains code that attempts to update a `sync.Map` from a goroutine.  Without proper synchronization, the result of loading the value from the map is non-deterministic.

## Solution

The `bugSolution.go` file illustrates a corrected approach using a mutex to protect the concurrent access to the `sync.Map`, ensuring consistent results.