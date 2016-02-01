package wald

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBernoulliSampleMLE(t *testing.T) {
	tests := []struct {
		in  BernoulliSample
		out float64
	}{
		{BernoulliSample{}, math.NaN()},
		{BernoulliSample{-1, 0}, math.NaN()},
		{BernoulliSample{-1, -1}, math.NaN()},
		{BernoulliSample{10, 0}, math.NaN()},
		{BernoulliSample{10, -1}, math.NaN()},
		{BernoulliSample{10, 2}, math.NaN()},
		{BernoulliSample{10, 20}, 0.5},
		{BernoulliSample{0, 400}, 0.0},
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

func TestBernoulliSampleVariance(t *testing.T) {
	tests := []struct {
		in  BernoulliSample
		out float64
	}{
		{BernoulliSample{}, math.NaN()},
		{BernoulliSample{-1, 0}, math.NaN()},
		{BernoulliSample{-1, -1}, math.NaN()},
		{BernoulliSample{10, 0}, math.NaN()},
		{BernoulliSample{10, -1}, math.NaN()},
		{BernoulliSample{10, 2}, math.NaN()},
		{BernoulliSample{0, 400}, 0.0},
		{BernoulliSample{10, 20}, 0.0125},
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
