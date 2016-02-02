package waldo

import "math"

// Sample represents data drawn from some distribution.  To compute
// the Wald statistics we need to have a point estimator function
// (e.g., the maximum likelihood estimator (MLE))
// as well as the sampling distribution's variance.  Recall
// that the sampling distribution is defined as the distribution of
// the point estimator.
//
// The estimator in question should be
// asymptotically normal, which is to say that the difference between
// the estimator (as a random variable of the data size) and the parameter
// being estimated over the standard error of the estimator converges
// in distribution to a standard normal distribution.
type Sample interface {
	Estimator() float64
	Variance() float64
}

// sample converts a pair (param estimate, variance) into
// a Sample implementation.
type sample struct {
	mle      float64
	variance float64
}

func (s sample) Estimator() float64 { return s.mle }
func (s sample) Variance() float64  { return s.variance }

// NewSample converts a sample parameter estimate and variance into a
// struct that implements the Sample interface.
func NewSample(estimate, variance float64) Sample {
	return sample{mle: estimate, variance: variance}
}

// StandardError computes an estimate for the standard error
// of a point estimator, as encoded in a Sample.
// The standard error is the standard deviation of the estimator's distribution.
// SInce the variance of this distribution is estimated, hence the overall
// calculation itself is an estimate.
func StandardError(s Sample) float64 {
	return math.Pow(s.Variance(), 0.5)
}
