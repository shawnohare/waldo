package wald

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairedComparisonMLE(t *testing.T) {
	tests := []struct {
		in  PairedComparison
		out float64
	}{
		// case NaN
		{
			in: PairedComparison{
				X: BernoulliSample{},
				Y: BernoulliSample{},
			},
			out: math.NaN(),
		},
		{
			in: PairedComparison{
				X: BernoulliSample{},
				Y: BernoulliSample{10, 20},
			},
			out: math.NaN(),
		},
		{
			in: PairedComparison{
				X: BernoulliSample{},
				Y: BernoulliSample{30, 20},
			},
			out: math.NaN(),
		},
		{
			in: PairedComparison{
				X: BernoulliSample{-1, 2},
				Y: BernoulliSample{10, 20},
			},
			out: math.NaN(),
		},
		{
			in: PairedComparison{
				X: BernoulliSample{10, 10},
				Y: BernoulliSample{5, 10},
			},
			out: 0.5,
		},
		{
			in: PairedComparison{
				X: BernoulliSample{10, 20},
				Y: BernoulliSample{10, 10},
			},
			out: -0.5,
		},
	}

	for _, tt := range tests {
		actual := tt.in.MLE()
		if math.IsNaN(tt.out) {
			assert.True(t, math.IsNaN(actual))
		} else {
			assert.Equal(t, tt.out, actual)
		}
	}
}

func TestPairedComparisonVariance(t *testing.T) {
	tests := []struct {
		in  PairedComparison
		out float64
	}{
		// case NaN
		{
			in: PairedComparison{
				X: BernoulliSample{},
				Y: BernoulliSample{},
			},
			out: math.NaN(),
		},
		// case 1
		{
			in: PairedComparison{
				X: BernoulliSample{},
				Y: BernoulliSample{},
			},
			out: math.NaN(),
		},
		// case 2
		{
			in: PairedComparison{
				X: BernoulliSample{},
				Y: BernoulliSample{},
			},
			out: math.NaN(),
		},
	}

	for _, tt := range tests {
		actual := tt.in.Variance()
		if math.IsNaN(tt.out) {
			assert.True(t, math.IsNaN(actual))
		} else {
			assert.Equal(t, tt.out, actual)
		}
	}
}
