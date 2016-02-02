package waldo

// PairedComparison tests whether two parameter estimates are equal.
// A typical use-case is to determine whether two Bernoulli samples
// are drawn from the same distribution.
type PairedComparison struct {
	X Sample
	Y Sample
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
