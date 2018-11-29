// Package scrabble implements functions relating to the board game Scrabble
package scrabble

import "unicode"

// Score returns the total score of a provided english word
func Score(word string) int {
  var total int

  for _, letter := range word {
    total += Points(letter)
  }

  return total
}

// Points returns the score of a single letter in the English edition
// A, E, I, O, U, L, N, R, S, T       1
// D, G                               2
// B, C, M, P                         3
// F, H, V, W, Y                      4
// K                                  5
// J, X                               8
// Q, Z                               10
func Points(letter rune) int {
  switch unicode.ToLower(letter) {
  case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
    return 1
  case 'd', 'g':
    return 2
  case 'b', 'c', 'm', 'p':
    return 3
  case 'f', 'h', 'v', 'w', 'y':
    return 4
  case 'k':
    return 5
  case 'j', 'x':
    return 8
  case 'q', 'z':
    return 10
  default:
    return 0
  }
}
