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

func ExampleWald() {
	data := BernoulliSample{Successes: 75, Count: 100}
	wald := Wald{Size: 0.05, Null: 0.5}
	fmt.Printf("%#v", wald.Test(data))
	// waldo.Result{ConfidenceInterval:[]float64{0.8168459395887369, 0.9831540604112632}, ConfidenceLevel:0.95, Power:0.9999999999999594, PValue:4.176224939330939e-21, RejectNull:true, Statistic:9.428090415820636}fmt.Println(r)
}
