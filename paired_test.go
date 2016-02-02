package waldo

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairedComparisonMLE(t *testing.T) {
	tests := []struct {
		in  PairedSample
		out float64
	}{
		// case NaN
		{
			in: PairedSample{
				X: BernoulliSample{},
				Y: BernoulliSample{},
			},
			out: math.NaN(),
		},
		{
			in: PairedSample{
				X: BernoulliSample{},
				Y: BernoulliSample{10, 20},
			},
			out: math.NaN(),
		},
		{
			in: PairedSample{
				X: BernoulliSample{},
				Y: BernoulliSample{30, 20},
			},
			out: math.NaN(),
		},
		{
			in: PairedSample{
				X: BernoulliSample{-1, 2},
				Y: BernoulliSample{10, 20},
			},
			out: math.NaN(),
		},
		{
			in: PairedSample{
				X: BernoulliSample{10, 10},
				Y: BernoulliSample{5, 10},
			},
			out: 0.5,
		},
		{
			in: PairedSample{
				X: BernoulliSample{10, 20},
				Y: BernoulliSample{10, 10},
			},
			out: -0.5,
		},
	}

	for _, tt := range tests {
		actual := tt.in.Estimator()
		if math.IsNaN(tt.out) {
			assert.True(t, math.IsNaN(actual))
		} else {
			assert.Equal(t, tt.out, actual)
		}
	}
}

func TestPairedComparisonVariance(t *testing.T) {
	tests := []struct {
		in  PairedSample
		out float64
	}{
		// case NaN
		{
			in: PairedSample{
				X: BernoulliSample{},
				Y: BernoulliSample{},
			},
			out: math.NaN(),
		},
		// case 1
		{
			in: PairedSample{
				X: BernoulliSample{},
				Y: BernoulliSample{},
			},
			out: math.NaN(),
		},
		// case 2
		{
			in: PairedSample{
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

func TestPairedComparisonTest(t *testing.T) {
	size := 0.05
	X := BernoulliSample{103, 200}
	Y := BernoulliSample{110, 200}
	ps := PairedSample{X: X, Y: Y}
	wald := Wald{Size: size}
	assert.Equal(t, wald.Test(ps), PairedComparisonTest(X, Y, size))
}

func TestPairedComparisonTestMethod(t *testing.T) {
	size := 0.05
	s := PairedSample{
		X: BernoulliSample{103, 200},
		Y: BernoulliSample{110, 200},
	}
	pc := PairedComparison{Size: size}
	wald := Wald{Size: size}
	assert.Equal(t, wald.Test(s), pc.Test(s))
}

func ExamplePairedComparison_WaldTest() {
	size := 0.05
	s := PairedSample{
		X: BernoulliSample{103, 200},
		Y: BernoulliSample{110, 200},
	}
	pc := PairedComparison{Size: size}
	fmt.Printf("%#v", pc.Test(s))
	// Output: waldo.Result{ConfidenceInterval:[]float64{-0.1327305906069241, 0.06273059060692406}, ConfidenceLevel:0.95, Power:0.10807314041617873, PValue:0.482731935542819, RejectNull:false, Statistic:-0.7019153324868983}
}

func ExamplePairedComparisonTest() {
	size := 0.05
	X := BernoulliSample{103, 200}
	Y := BernoulliSample{110, 200}
	fmt.Printf("%#v", PairedComparisonTest(X, Y, size))
	// Output: waldo.Result{ConfidenceInterval:[]float64{-0.1327305906069241, 0.06273059060692406}, ConfidenceLevel:0.95, Power:0.10807314041617873, PValue:0.482731935542819, RejectNull:false, Statistic:-0.7019153324868983}
}
