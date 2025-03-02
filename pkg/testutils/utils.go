// Copyright 2025 César Nieto Sánchez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package testutils

import (
	"fmt"
)

// CompareSlices compares two slices of nD points.
// Arguments:
//   - points ([][]float64): The points to compare.
//   - expected ([][]float64): The expected points.
//
// Returns:
//   - bool: True if the points are equal, false otherwise.
//   - error: An error if the points are not equal.
func CompareSlices(points [][]float64, expected [][]float64) (bool, error) {

	if len(points) != len(expected) {
		return false, fmt.Errorf("Expected %v points, but got %v", len(expected), len(points))
	}

	errorMsg := ""
	for i, point := range points {
		for j := range point {
			if points[i][j] != expected[i][j] {
				errorMsg += fmt.Sprintf("\nAt coordinates (%v,%v)  expected value %v, but got %v", i, j, expected[i][j], points[i][j])
			}
		}
	}
	if len(errorMsg) > 0 {
		return false, errors.New(errorMsg)
	}

	return true, nil

}
