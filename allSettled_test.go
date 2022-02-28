package asynks

import (
	"testing"

	"github.com/golang-must/must"
)

func TestAllSettled(t *testing.T) {
	tasks := []TaskSettled{
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

	t.Run("data length equals task length", func(t *testing.T) {
		data := AllSettled(tasks)

		must := must.New(t)
		must.Equal(len(data), 3)
	})
}
