// Package accumulate implements functions for mapping collections
package accumulate

// Accumulate transforms a collection of strings into a second collection,
// with the supplied converter applied to each.
func Accumulate(xs []string, converter func(string) string) []string {
	length := len(xs)
	ys := make([]string, length)

	for i := 0; i < length; i++ {
		ys[i] = converter(xs[i])
	}

	return ys
}
