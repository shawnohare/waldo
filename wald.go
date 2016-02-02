package waldo

import "math"

// ScalarParameterHypothesisTest represents an abstract hypothesis test for
// a scalar parameter, of which the Wald test is one.
// This type is introduced in case the waldo package becomes part of
// a larger class of tests.
// type ScalarParameterHypothesisTest interface {
// 	Test(s Sample) Result
// }

// Wald is a simple two-sided statistical hypothesis test
// for a scalar parameter theta. It tests
// H_0: theta = null vs. H_1: theta != null,
// under the assumption that the estimator for theta is asymptotically
// normal. Asymptotic normality is met by Bernoulli trials with the
// maximum likelihood estimator, for instance.
// A Wald test with Size S has a confidence level of 1 - S.
type Wald struct {
	Size float64
	Null float64
	zval float64
}

// Results is a container for the result of performing a Wald
// test over some sample of data.
type Result struct {
	// Confidence interval for the sample parameter estimate computed at
	// the same level as the size of the test.
	ConfidenceInterval []float64
	// The confidence level is 1- the size of the Wald test that produced the result.
	ConfidenceLevel float64
	Power           float64
	PValue          float64
	RejectNull      bool
	Statistic       float64
}

// Compute the z_{alpha/2} value associated wit the test.
func (t *Wald) z() float64 {
	if t.zval == 0 {
		t.zval = zalpha2(t.Size)
	}
	return t.zval
}

// Statistic computes the test statistic for the sample.
func (t Wald) Statistic(s Sample) float64 {
	return (s.Estimator() - t.Null) / StandardError(s)
}

func (t Wald) PValue(s Sample) float64 {
	return 2 * stdNormal.Cdf(-math.Abs(t.Statistic(s)))
}

func (t Wald) ConfidenceInterval(s Sample) []float64 {
	estimate := s.Estimator()
	stdError := StandardError(s)
	epsilon := stdError * t.z()
	return []float64{estimate - epsilon, estimate + epsilon}
}

// Power function estimate. The power
// is the probability of correctly rejecting the null hypothesis.
func (t Wald) Power(s Sample) float64 {
	estimate := s.Estimator()
	stdError := StandardError(s)
	z := t.z()
	x := (t.Null - estimate) / stdError
	return 1 - stdNormal.Cdf(x+z) + stdNormal.Cdf(x-z)
}

func (t Wald) RejectNull(s Sample) bool {
	return math.Abs(t.Statistic(s)) > t.z()
}

// Perform a Wald test on a sample of data.
func (t Wald) Test(s Sample) Result {
	cs := NewSample(s.Estimator(), s.Variance()) // computed sample
	return Result{
		Statistic:          t.Statistic(cs),
		ConfidenceInterval: t.ConfidenceInterval(cs),
		ConfidenceLevel:    1 - t.Size,
		Power:              t.Power(cs),
		PValue:             t.PValue(cs),
		RejectNull:         t.RejectNull(cs),
	}
}
