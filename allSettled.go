package asynks

import (
	"sync"
)

type TaskSettled func() (interface{}, error)
type dataSettled struct {
	Status string
	Data   interface{}
}

func AllSettled(tasks []TaskSettled) []dataSettled {
	wg := &sync.WaitGroup{}
	results := make([]dataSettled, 0)

	for _, task := range tasks {
		wg.Add(1)
		go func(t TaskSettled) {
			defer wg.Done()
			var result dataSettled
			data, err := t()
			if err != nil {
				result.Status = "error"
				result.Data = err.Error()
			} else {
				result.Status = "ok"
				result.Data = data
			}
			results = append(results, result)
		}(task)
	}

	wg.Wait()
	return results
}
