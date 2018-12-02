// Package diffsquares implements functions about sums and squares
package diffsquares

// SquareOfSum returns the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the square of each of the first N natural numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns the SquareOfSum less the SumOfSquares
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
