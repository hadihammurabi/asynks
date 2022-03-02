package asynks

import (
	"sync"
)

type TaskAll func() (interface{}, error)

// All runs all tasks in concurrent way
func All(tasks []TaskAll) ([]interface{}, error) {
	wg := &sync.WaitGroup{}
	results := make([]interface{}, 0)
	var gerr error
	var gerrOnce sync.Once

	for _, task := range tasks {
		wg.Add(1)
		go func(t TaskAll, eo *sync.Once) {
			defer wg.Done()
			data, err := t()
			if err != nil {
				gerrOnce.Do(func() {
					gerr = err
				})
			} else {
				results = append(results, data)
			}
		}(task, &gerrOnce)
	}

	wg.Wait()
	if gerr != nil {
		return nil, gerr
	}

	return results, nil
}
