package asynks

type TaskRace func() (interface{}, error)

func Race(tasks []TaskRace) (interface{}, error) {
	result := make(chan interface{})
	gerr := make(chan error)

	for _, task := range tasks {
		go func(t TaskRace, re chan<- interface{}, ge chan<- error) {
			data, err := t()
			if err != nil {
				ge <- err
			} else {
				re <- data
			}
		}(task, result, gerr)
	}

	select {
	case err := <-gerr:
		return nil, err
	case data := <-result:
		return data, nil
	}
}
