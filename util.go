// github.com/benwebber/bitboard

// Copyright 2014 Ben Webber <benjamin.webber@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package bitboard implements 8x8 bitboards for games like chess, checkers,
// Reversi, and Othello.
package chess

import (
	"strconv"
)

var (
	// Number of columns
	files   = 8
	symbols = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
)

// Convert coordinates in algebraic notation to an integer bit position.
func AlgebraicToBit(p string) int {
	x, y := AlgebraicToCartesian(p)
	return CartesianToBit(x, y)
}

// Convert coordinates in algebraic notation to Cartesian coordinates.
func AlgebraicToCartesian(p string) (int, int) {
	symbols := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var x int
	for i, v := range symbols {
		if string(p[0]) == v {
			x = i
		}
	}
	y, _ := strconv.Atoi(string(p[1]))
	return x, (y - 1)
}

// Convert Cartesian coordinates to an integer bit position.
func CartesianToBit(x int, y int) int {
	bit := y*files + x
	return bit
}
