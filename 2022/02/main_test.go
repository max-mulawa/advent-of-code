package main

import (
	"github.com/stretchr/testify/require"
	"testing"
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
			require.Equal(t, 1, 1)
		})
	}
}
