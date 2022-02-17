package controller

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smmd/academy-go-q42021/wpool"
)

func Test_WorkerRequest(t *testing.T) {
	testCases := []struct {
		name           string
		wtype          string
		items          int
		itemsPerWorker int
		result         wpool.Request
		hasError       bool
		err            error
	}{
		{
			"validate type odd",
			"odd",
			10,
			5,
			wpool.Request{
				TypeOfJob:      "odd",
				NumberOfItems:  10,
				ItemsPerWorker: 5,
			},
			false,
			nil,
		},
		{
			"validate type even",
			"even",
			20,
			5,
			wpool.Request{
				TypeOfJob:      "even",
				NumberOfItems:  20,
				ItemsPerWorker: 5,
			},
			false,
			nil,
		},
		{
			"invalidate type",
			"wrong",
			10,
			5,
			wpool.Request{
				TypeOfJob:      "wrong",
				NumberOfItems:  10,
				ItemsPerWorker: 5,
			},
			true,
			errors.New("invalid request: Key: 'Request.TypeOfJob' Error:Field validation for 'TypeOfJob' failed on the 'enum' tag"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := workerRequest(tc.wtype, tc.items, tc.itemsPerWorker)

			assert.Equal(t, tc.result, actual)

			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}
