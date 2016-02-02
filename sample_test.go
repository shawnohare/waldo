package waldo

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardError(t *testing.T) {
	tests := []struct {
		in  Sample
		out float64
	}{
		{sample{variance: -1}, math.NaN()},
		{NewSample(1, 4), 2.0},
	}

	for _, tt := range tests {
		actual := StandardError(tt.in)
		if math.IsNaN(tt.out) {
			assert.True(t, math.IsNaN(actual))
		} else {
			assert.Equal(t, tt.out, actual)
		}
	}
}
