package waldo

import (
	"math"
)

// BernoulliSample represents an IID sample drawn from a Bernoulli
// distribution with some unknown success probability p.
// Some examples include a series of weighted coin flips or clickthrough data
// for a campaign.  An BernoulliSample instance can be created either from
// a slice of data or a count of successes together with the number of trials.
type BernoulliSample struct {
	Successes int
	Trials    int
}

// MLE estimate for the underlying success probability.
// MLE computes the maximum likelihood estimator (MLE) estimate
// from the sample. If X := sum_{i=1}^n X_i, where X_i are
// iid from a Bernoulli distribution with unknown parameter p,
// then the MLE (as a function) is simply (1/m)X, where m is the size
// of the sample.
func (s BernoulliSample) Estimator() float64 {
	if s.Trials <= 0 || s.Successes < 0 || s.Successes > s.Trials {
		return math.NaN()
	}
	return float64(s.Successes) / float64(s.Trials)
}

func (s BernoulliSample) Variance() float64 {
	p := s.Estimator()
	return (p * (1 - p)) / float64(s.Trials)
}
