package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	prev := randInt
	defer func() { randInt = prev }()

	calls := 0
	randInt = func() int {
		calls++
		return 42
	}

	const n = 5
	got := generateRandomElements(n)

	assert.Len(t, got, n)

	for _, v := range got {
		assert.Equal(t, 42, v)
	}

	assert.Equal(t, n, calls, "randInt должен быть вызван n раз")
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"nil слайс", nil, 0},
		{"обычный слайс", []int{1, 2, 3, 10, 7}, 10},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := maximum(tc.in)
			assert.Equal(t, tc.want, got, "тест кейс: %s", tc.name)
		})
	}
}
