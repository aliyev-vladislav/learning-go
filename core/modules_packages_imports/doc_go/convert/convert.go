// Package convert provides various utilities to
// make it easy to convert money from one currency to another.
package convert

// Money represents the combination of an amount of money
// and the currency the money is in.
// More information on exchange rates can be found at [Investopedia].
type Money struct {
	Value    int
	Currency string
}

// Convert converts the value of one currency to another.
//
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns
// the converted currency and any errors encountered from unknown or unconvertible
// currencies.
//
// If an error is returned, the Money instance is set to the zero value.
//
// Supported currencies are:
//   - USD - US Dollar
//   - CAD - Canadian Dollar
//   - EUR - Euro
//   - INR - Indian Rupee
//
// More information on exchange rates can be found at [Investopedia].
//
// [Investopedia]: https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
	// ...
	return Money{}, nil
}
