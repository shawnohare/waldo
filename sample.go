package wald

import "math"

// Sample represents a sample drawn from some distribution.  To compute
// the Wald statistics we need to have a maximum-likelihood estimator
// function as well as the variance.
type Sample interface {
	// Maximum likelihood estimator evaluated over the sample data.
	MLE() float64
	// Variance of the sampling distribution.
	Variance() float64
}

// sample converts a pair (param estimate, variance) into
// a Sample implementation.
type sample struct {
	mle      float64
	variance float64
}

func (s sample) MLE() float64      { return s.mle }
func (s sample) Variance() float64 { return s.variance }

// NewSample converts a sample parameter estimate and variance into a
// struct that implements the Sample interface.
func NewSample(estimate, variance float64) Sample {
	return sample{mle: estimate, variance: variance}
}

// StandardError estimates the standard error of a sample.
// The standard error is the standard deviation of the MLE's distribution.
// The variance of this distribution is estimated, hence the overall
// calculation itself is an estimate.
func StandardError(s Sample) float64 {
	return math.Pow(s.Variance(), 0.5)
}
