package wpool

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smmd/academy-go-q42021/model"
)

var pokeMonstersToGenerate = []string{
	"bulbasaur",
	"ivysaur",
	"venusaur",
	"charmander",
	"charmeleon",
	"charizard",
	"squirtle",
	"wartortle",
	"blastoise",
	"caterpie",
	"metapod",
	"butterfree",
}

var evenPokeMonstersExpected = []*model.Pokemon{
	{
		"1",
		"bulbasaur",
	},
	{
		"3",
		"venusaur",
	},
	{
		"5",
		"charmeleon",
	},
	{
		"7",
		"squirtle",
	},
	{
		"9",
		"blastoise",
	},
}

var oddPokeMonstersExpected = []*model.Pokemon{
	{
		"2",
		"ivysaur",
	},
	{
		"4",
		"charmander",
	},
	{
		"6",
		"charizard",
	},
	{
		"8",
		"wartortle",
	},
}

func Test_Worker(t *testing.T) {
	testCases := []struct {
		name       string
		jobsInput  []string
		expected   []*model.Pokemon
		conditions *Conditions
		err        error
	}{
		{
			"append results even IDs",
			pokeMonstersToGenerate,
			evenPokeMonstersExpected,
			generateConditions(false, 5, 5),
			nil,
		},
		{
			"append results odd IDs",
			pokeMonstersToGenerate,
			oddPokeMonstersExpected,
			generateConditions(true, 4, 4),
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testResult := make(chan Result)
			testJobs := make(chan []string, len(tc.jobsInput))

			var wg sync.WaitGroup
			wg.Add(1)

			for k, v := range tc.jobsInput {
				value := []string{strconv.Itoa(k + 1), v}
				testJobs <- value
			}
			close(testJobs)

			go func() {
				defer wg.Done()
				worker(testJobs, testResult, tc.conditions)
			}()

			go func() {
				wg.Wait()
				close(testResult)
			}()

			i := 0
			for r := range testResult {
				assert.Equal(t, tc.expected[i], r.Result)
				assert.NoError(t, r.Err)
				i++
			}
		})
	}
}

func generateConditions(isOdd bool, limitPerWorker int, maxItems int) *Conditions {
	c := new(Conditions)
	c.isOdd = isOdd
	c.limitPerWorker = limitPerWorker
	c.maxItems = maxItems

	return c
}
