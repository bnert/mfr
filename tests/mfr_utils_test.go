package tests

import (
	"testing"

	mfru "github.com/bnert/mfr/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	assert.True(t, mfru.Contains([]int{0, 2, 4}, 2))
	assert.False(t, mfru.Contains([]int{0, 2, 4}, 3))
}

func Test_All(t *testing.T) {
	v := []int{0, 2, 4, 6, 8}
	assert.True(t, mfru.IsAll(v, func(item int) bool {
		return item%2 == 0
	}))
	assert.False(t, mfru.IsAll(v, func(item int) bool {
		return item%3 == 0
	}))
}

func Test_Some(t *testing.T) {
	v := []int{0, 1, 2, 3, 4}

	assert.True(t, mfru.IsSome(v, func(item int) bool {
		return item%2 == 0
	}))

	assert.True(t, mfru.IsSome(v, func(item int) bool {
		return item%3 == 0
	}))

	assert.False(t, mfru.IsSome(v, func(item int) bool {
		return item > 10
	}))
}

func Test_None(t *testing.T) {
	v := []int{0, 1, 2, 3, 4}
	assert.True(t, mfru.IsNone(v, func(item int) bool {
		return item == 100
	}))

	assert.False(t, mfru.IsNone(v, func(item int) bool {
		return item == 4
	}))
}

func Test_Reverse(t *testing.T) {
	v := []int{0, 1, 2, 3, 4}
	r := []int{4, 3, 2, 1, 0}
	assert.Equal(t, r, mfru.Reverse[int](v))
	assert.Equal(t, v, mfru.Reverse[int](mfru.Reverse[int](v)))
}

type DStruct struct {
	Item string
}

func Test_Distinct(t *testing.T) {
	v1 := []int{0, 1, 2, 3, 4}
	r1 := v1
	assert.Equal(t, r1, mfru.Distinct(v1))
	assert.True(t, mfru.IsDistinct(v1))

	v2 := []int{0, 1, 1, 2, 3}
	r2 := []int{0, 1, 2, 3}
	assert.Equal(t, r2, mfru.Distinct(v2))
	assert.False(t, mfru.IsDistinct(v2))

	v3 := []int{0, 0, 0, 0, 0}
	r3 := []int{0}
	assert.Equal(t, r3, mfru.Distinct(v3))
	assert.False(t, mfru.IsDistinct(v3))

	v4 := []DStruct{
		DStruct{"hello"},
		DStruct{"world"},
	}
	assert.Equal(t, v4, mfru.Distinct(v4))

	v5 := []DStruct{
		DStruct{"hello"},
		DStruct{"world"},
		DStruct{"hello"},
	}
	r5 := []DStruct{
		DStruct{"hello"},
		DStruct{"world"},
	}
	assert.Equal(t, r5, mfru.Distinct(v5))
}
