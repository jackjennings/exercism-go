// Package accumulate implements functions for mapping collections
package accumulate

// Accumulate transforms a collection of strings into a second collection,
// with the supplied converter applied to each.
func Accumulate(xs []string, converter func(string) string) []string {
	ys := make([]string, len(xs))

	for i, x := range xs {
		ys[i] = converter(x)
	}

	return ys
}
