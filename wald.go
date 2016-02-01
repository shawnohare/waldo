package waldo

import "math"

// TODO
// - Refactor WaldTest to

// Test is a size alpha simple two-sided statistical hypothesis test
// for a scalar parameter theta. It tests
// H_0: theta = null vs. H_1: theta != null,
// under the assumption that the estimator for theta is asymptotically
// normal. Asymptotic normality is met by Bernoulli trials with the
// maximum likelihood estimator, for instance.
type Test struct {
	Alpha float64 // Size of the test.
	Null  float64 // Null value for the test.
	zval  float64 // z_{alpha/2} value
}

// TestResults is a container for the result of performing a Wald
// test over some sample of data.
type TestResult struct {
	ConfidenceInterval []float64
	Power              float64
	PValue             float64
	RejectNull         bool
	Statistic          float64
}

// Compute the z_{alpha/2} value associated wit the test.
func (t *Test) z() float64 {
	if t.zval != 0 {
		return t.zval
	}
	return zalpha2(t.Alpha)
}

// Statistic computes the test statistic for the sample.
func (t Test) Statistic(s Sample) float64 {
	return (s.MLE() - t.Null) / StandardError(s)
}

func (t Test) PValue(s Sample) float64 {
	return 2 * stdNormal.Cdf(-math.Abs(t.Statistic(s)))
}

func (t Test) ConfidenceInterval(s Sample) []float64 {
	estimate := s.MLE()
	stdError := StandardError(s)
	epsilon := stdError * t.z()
	return []float64{estimate - epsilon, estimate + epsilon}
}

// Power function for the size alpha Wald test.  The value returned
// is the measure of the confidence interval computed from the estimate.
func (t Test) Power(s Sample) float64 {
	estimate := s.MLE()
	stdError := StandardError(s)
	z := t.z()
	x := (t.Null - estimate) / stdError
	return 1 - stdNormal.Cdf(x+z) + stdNormal.Cdf(x-z)
}

func (t Test) RejectNull(s Sample) bool {
	return math.Abs(t.Statistic(s)) > t.z()
}

// Perform a Wald test on a sample of data.
func (t Test) Perform(s Sample) TestResult {
	cs := NewSample(s.MLE(), s.Variance()) // computed sample
	return TestResult{
		Statistic:          t.Statistic(cs),
		ConfidenceInterval: t.ConfidenceInterval(cs),
		Power:              t.Power(cs),
		PValue:             t.PValue(cs),
		RejectNull:         t.RejectNull(cs),
	}
}
