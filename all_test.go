package asynks

import (
	"testing"

	"github.com/golang-must/must"
)

func TestAll(t *testing.T) {
	tasks := []TaskAll{
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

	// tasksWithErrors := []TaskAll{
	// 	func() (interface{}, error) {
	// 		return nil, errors.New("don't cry")
	// 	},
	// 	func() (interface{}, error) {
	// 		return nil, errors.New("don't cry")
	// 	},
	// 	func() (interface{}, error) {
	// 		return nil, errors.New("don't cry")
	// 	},
	// }

	t.Run("data length equals task length", func(t *testing.T) {
		data := All(tasks)

		must := must.New(t)
		must.Equal(3, len(data))
	})

	// t.Run("task error produce error", func(t *testing.T) {
	// 	must := must.New(t)
	// 	defer func() {
	// 		err := recover()
	// 		must.NotNil(err)
	// 	}()

	// 	data := All(tasksWithErrors)
	// 	must.Nil(data)
	// })
}
