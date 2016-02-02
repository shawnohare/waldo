package waldo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunczalpha2(t *testing.T) {
	tests := []struct {
		in  float64
		out float64
	}{
		{0.05, 1.95996},
		{0.01, 2.57583},
		{0.0, stdNormal.Ppf(1.0)},
		{0.001, stdNormal.Ppf(1.0 - (.001 / 2.0))},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.out, zalpha2(tt.in))
	}
}
