// Package diffsquares implements functions about sums and squares
package diffsquares

// SquareOfSum returns the square of the sum of the first N natural numbers
func SquareOfSum(input int) int {
  var output int

  for i := input; i > 0; i-- {
    output += i
  }

  return output * output
}

// SumOfSquares returns the sum of the square of each of the first N natural numbers
func SumOfSquares(input int) int {
  var output int

  for i := input; i > 0; i-- {
    output += i * i
  }

  return output
}

// Difference returns the SquareOfSum less the SumOfSquares
func Difference(input int) int {
  return SquareOfSum(input) - SumOfSquares(input)
}
