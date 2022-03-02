# asynks
Run asynchronous tasks in Go.

> Inspired by Promise methods in JS

# Features
* Run functions asynchronously
* Clear API
* Remove goroutine complexity

# Example
```go
package main

import (
	"fmt"

	// Import asynks
	"github.com/hadihammurabi/asynks"
)

func main() {
	// Use asynks.All to run all functions asynchronously
	results, err := asynks.All(
		func() (interface{}, error) {
			return 1, nil
		},
		func() (interface{}, error) {
			return 2, nil
		},
		func() (interface{}, error) {
			return 3, nil
		},
	)

	// Panic on error and print the results
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}
```
