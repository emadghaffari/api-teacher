package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	tests := []struct {
		step string
		len  int
		min  int
		max  int
	}{
		{
			step: "a",
			len:  10,
			min:  1,
			max:  10,
		},
		{
			step: "b",
			len:  100,
			min:  10,
			max:  1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {

			rn := Slice(tc.len, tc.min, tc.max)
			assert.Equal(t, tc.len, len(rn))
		})
	}
}

func TestRand(t *testing.T) {
	tests := []struct {
		step string
		min  int
		max  int
	}{
		{
			step: "a",
			min:  1,
			max:  10,
		},
		{
			step: "b",
			min:  10,
			max:  1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {

			rn := Rand(tc.min, tc.max)
			if rn > tc.min && rn < tc.max {
				assert.True(t, true)
			}
		})
	}
}
