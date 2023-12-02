package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	for _, tc := range []struct {
		l string
	}{
		{
			l: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		},
	} {
		t.Run(tc.l, func(t *testing.T) {
			g := NewGame(tc.l)
			require.Equal(t, g.b, 6)
			require.Equal(t, g.no, 1)
			require.True(t, g.valid())
		})
	}
}
