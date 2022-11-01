package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFloot(t *testing.T) {
	for _, tc := range []struct {
		input    string
		floor    int
		basement int
	}{
		{
			input:    "(())",
			floor:    0,
			basement: 0,
		},
		{
			input:    "()()",
			floor:    0,
			basement: 0,
		},
		{
			input:    "(((",
			floor:    3,
			basement: 0,
		},
		{
			input:    "(()(()(",
			floor:    3,
			basement: 0,
		},
		{
			input:    "))(((((",
			floor:    3,
			basement: 1,
		},
		{
			input:    "()())",
			floor:    -1,
			basement: 5,
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			f, basementPos := floor([]byte(tc.input))
			require.Equal(t, tc.floor, f)
			require.Equal(t, tc.basement, basementPos)
		})
	}
}
