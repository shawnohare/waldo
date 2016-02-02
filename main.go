// Package waldo performs basic statistical hypothesis testing for scalar
// parameter estimates using the Wald test.
package waldo

import (
	"github.com/chobie/go-gaussian"
)

// stdNormal is a standard normal distribution.
var stdNormal = gaussian.NewGaussian(0, 1)

// zalpha2 computes the standard normal's inverse CDF of 1 - (alpha/2).
// This computes the value z such that the probability that Z < z
// is 1 - (alpha/2), where Z is a random variable distributed according
// Alternatively, P(Z > z_{\alpha/2}) = \alpha/2.
// This value appears frequently when we wish to compute
// (1 - \alpha) normal-based confidence intervals or perform statistical hypothesis
// testing. For example, the normal based 95% confidence interval for a parameter
// estimate hat(p) is the epsilon ball B(hat(p), hat(se)*zalpha2(1-0.95)).
func zalpha2(alpha float64) float64 {
	// Define some hard-coded values for frequently used alpha values.
	// These correspond to the 99.5, 99, ..., 97.5 percentile point values
	// of the standard normal distribution.
	hardcoded := map[float64]float64{
		0.01: 2.57583,
		0.02: 2.32635,
		0.03: 2.17009,
		0.04: 2.0537,
		0.05: 1.95996,
	}

	if z, ok := hardcoded[alpha]; ok {
		return z
	}

	// Ppf is the percentile point function, i.e., inverse CDF.
	return stdNormal.Ppf(1.0 - (alpha / 2.0))
}
