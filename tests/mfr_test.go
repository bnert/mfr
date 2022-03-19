package tests

import (
	"strconv"
	"testing"

	mfr "github.com/bnert/mfr"
	mfru "github.com/bnert/mfr/utils"
	"github.com/stretchr/testify/assert"
)

func Test_SliceMap(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}

	s1 := mfr.Map[int, string](s, func(ctx mfr.Ctx[int]) string {
		return strconv.Itoa(ctx.Item * 2)
	})

	assert.Equal(t, []string{"0", "2", "4", "6", "8"}, s1)
}

func Test_SliceMapFilterReduce(t *testing.T) {
	s := []string{"Sev'ral", "omit", "Timez", "4", "me", "eva"}

	cleaned := mfr.Filter[string](s, func(ctx mfr.Ctx[string]) bool {
		return !mfru.Contains([]string{"omit", "me"}, ctx.Item)
	})

	wPeriods := mfr.Map[string, string](cleaned, func(ctx mfr.Ctx[string]) string {
		return ctx.Item + "."
	})

	allTogether := mfr.Reduce[string, string](wPeriods, "", func(ctx mfr.Ctx[string], acc string) string {
		if ctx.IsLast {
			return acc + ctx.Item
		}
		return acc + ctx.Item + " "
	})

	assert.Equal(t, "Sev'ral. Timez. 4. eva.", allTogether)
}

func Test_MapEntries(t *testing.T) {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	targetResult := []mfr.Entry[string, int]{
		mfr.Entry[string, int]{Key: "one", Value: 1},
		mfr.Entry[string, int]{Key: "two", Value: 2},
	}

	entries := mfr.EntriesFromMap[string, int](m)

	assert.ElementsMatch(t, targetResult, entries)
}

func Test_MapEntriesToMap(t *testing.T) {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	entries := mfr.EntriesFromMap[string, int](m)
	m1 := mfr.EntriesToMap[string, int](entries)

	assert.Equal(t, m, m1)
}
