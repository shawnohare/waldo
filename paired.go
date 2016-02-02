package waldo

// PairedSample represents data used to determine whether the parameter
// estimates for the random variables X and Y agree.
// A typical use-case is to determine whether two Bernoulli samples {X, Y}
// are drawn from the same distribution.
//
// As a random variable, a paired sample represents Z := X - Y.
type PairedSample struct {
	X Sample
	Y Sample
}

// PairedComparisonTest is a Wald Hypothesis test of the specified size
// for paired sample data.
type PairedComparison struct {
	Size float64
}

func (t PairedComparison) Test(pairedSample Sample) Result {
	return Wald{Size: t.Size}.Test(pairedSample)
}

// Estimator for the comparison estimator hat(p) := X.Estimator - Y.Estimator
func (s PairedSample) Estimator() float64 {
	return s.X.Estimator() - s.Y.Estimator()
}

// Variance for a PairedSample is the sum of the individual variances.
// This follows from the fact that the variance operator is additive over
// independent variables, which X and Y are presumed to be, and scalars
// factor out as their squares.  That is, if V is the variance operator, then
// V(X-Y) = V(X) + V(-Y) = V(X) + (-1)^{2}V(Y).
func (s PairedSample) Variance() float64 {
	// V(X - Y) = V(X) + V(-Y) = V(X) + (-1)^2*V(Y)
	return s.X.Variance() + s.Y.Variance()
}

// PairedComparisonTest performs a Wald comparison test of the specified
// size for the paired data (X, Y).  It wraps PairedComparison.Test.
func PairedComparisonTest(X Sample, Y Sample, size float64) Result {
	pc := PairedComparison{Size: size}
	ps := PairedSample{X: X, Y: Y}
	return pc.Test(ps)
}
