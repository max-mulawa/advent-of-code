package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	for _, tc := range []struct {
		desc string
	}{
		{
			desc: "",
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			l := "Card 186: 25  3 81 78 75 48 38 71 43 80 | 58 56 22 93 69  2  6 14 36 66 31 50 67 53 27 86 95 72 55 46 12 35 34 96 16"

			c := NewCard(l)

			require.Equal(t, 0, c.value())
			require.Equal(t, "", c.wining())
		})
	}
}
