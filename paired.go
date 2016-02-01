package wald

// PairedComparison tests whether two Bernoulli samples are drawn from
// the same underlying distribution, utilizing a Wald test.  If X is the
// number of successes in the first sample (of size m) and Y the same
// value for the second sample (of size n),
// then X ~ Binomial(m, p1) and Y ~ Binomial(n,p1)
// for unkown parameter's p1 and p2.
type PairedComparison struct {
	X BernoulliSample
	Y BernoulliSample
}

// MLE estimate of the comparison estimator hat(p) := X.MLE() - Y.MLE()
func (s PairedComparison) MLE() float64 {
	return s.X.MLE() - s.Y.MLE()
}

// StandardError estimates the standard deviation of the distribution
// to which the MLE belongs.
func (s PairedComparison) Variance() float64 {
	// V(X - Y) = V(X) + V(-Y) = V(X) + (-1)^2*V(Y)
	return s.X.Variance() + s.Y.Variance()
}

// FIXME can probably delete this.
// WaldStatistic associated to the comparison sample, using a Wald
// test of size alpha.
// func (s PairedComparison) WaldStatistic(alpha float64) *Statistic {
// 	p := s.MLE()
// 	se := StandardError(s)
// 	t := Test{Alpha: alpha}
// 	return t.Statistic(p, se)
// }
