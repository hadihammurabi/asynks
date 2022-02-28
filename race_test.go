package asynks

import (
	"errors"
	"testing"

	"github.com/golang-must/must"
)

func TestRace(t *testing.T) {
	tasks := []TaskRace{
		func() (interface{}, error) {
			return 1, nil
		},
		func() (interface{}, error) {
			return 2, nil
		},
		func() (interface{}, error) {
			return 3, nil
		},
	}

	tasksWithErrors := []TaskRace{
		func() (interface{}, error) {
			return nil, errors.New("don't cry")
		},
		func() (interface{}, error) {
			return nil, errors.New("don't cry")
		},
		func() (interface{}, error) {
			return nil, errors.New("don't cry")
		},
	}

	t.Run("data exists in returns", func(t *testing.T) {
		data, _ := Race(tasks)

		must := must.New(t)
		must.True(data == 1 || data == 2 || data == 3)
	})

	t.Run("task error produce error", func(t *testing.T) {
		data, err := Race(tasksWithErrors)

		must := must.New(t)
		must.Nil(data)
		must.NotNil(err)
	})
}
