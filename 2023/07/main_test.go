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
			//hand: K5JK5 with 5 score 490 bid, rank: 952
			c := &Cards{cards: "K5JK5"}
			require.Equal(t, FullHouse, c.score())
		})
	}
}
