package asynks

import (
	"sync"
)

type TaskAll func() (interface{}, error)

// All runs all tasks in concurrent way
// TODO: it panic when error happen, instead returns error
func All(tasks []TaskAll) []interface{} {
	wg := &sync.WaitGroup{}
	results := make([]interface{}, 0)

	for _, task := range tasks {
		wg.Add(1)
		go func(t TaskAll) {
			defer wg.Done()
			data, err := t()
			if err != nil {
				panic(err)
			}
			results = append(results, data)
		}(task)
	}

	wg.Wait()
	return results
}
