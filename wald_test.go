package waldo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaldZ(t *testing.T) {
	tests := []struct {
		in  Wald
		out float64
	}{
		{Wald{Size: 0.05}, 1.95996},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.out, tt.in.z())
		// Get cached value.
		assert.Equal(t, tt.out, tt.in.z())
	}
}

func TestWaldTest(t *testing.T) {
	tests := []struct {
		wald   Wald
		sample Sample
		result Result
	}{
		{
			wald:   Wald{Size: 0.05, Null: 0},
			sample: sample{mle: 1, variance: 1},
			result: Result{
				Statistic:          1.0,
				Power:              1 - stdNormal.Cdf(-1+1.95996) + stdNormal.Cdf(-1-1.95996),
				ConfidenceInterval: []float64{-0.95996, 2.95996},
				ConfidenceLevel:    .95,
				PValue:             2 * stdNormal.Cdf(-1.0),
				RejectNull:         false,
			},
		},
	}

	for _, tt := range tests {
		actual := tt.wald.Test(tt.sample)
		assert.Equal(t, tt.result.Statistic, actual.Statistic)
		assert.Equal(t, tt.result.Power, actual.Power)

		// Confidence interval
		assert.InEpsilon(t, tt.result.ConfidenceInterval[0], actual.ConfidenceInterval[0], 0.0001)
		assert.InEpsilon(t, tt.result.ConfidenceInterval[1], actual.ConfidenceInterval[1], 0.0001)

		assert.Equal(t, tt.result.PValue, actual.PValue)
		assert.Equal(t, tt.result.RejectNull, actual.RejectNull)
	}
}

func ExampleWaldTestOnBernoulliSample() {
	data := BernoulliSample{Successes: 35, Trials: 40}
	wald := Wald{Size: 0.05, Null: 0.5}
	fmt.Printf("%#v", wald.Test(data))
	// Output: waldo.Result{ConfidenceInterval:[]float64{0.7725112383996421, 0.9774887616003579}, ConfidenceLevel:0.95, Power:0.999999906295462, PValue:7.425001949510999e-13, RejectNull:true, Statistic:7.171371656006362}
}
