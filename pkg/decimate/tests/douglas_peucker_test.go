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
package tests

import (
	"fmt"
	"github.com/cenieto/decimate/pkg/geom2d"
	"testing"
)

// // TestValidateInputPointListOnePoint tests the input check of the ValidateInputPointList function.
// // It checks if the function panics when the input has only one element.
// //
// // Parameters:
// //   - t (*testing.T): A testing object used to run tests and check for failures.
// //
// // Returns:
// //   - None
// func TestValidateInputPointListOnePoint(t *testing.T) {
// 	points := [][]float64{
// 		{1.0, 2.0},
// 	}

// 	geometry := geom2d.NewEuclid()

// 	errorMsg := geometry.Decimate.ValidateInputPointList(points)
// 	if errorMsg != nil {
// 		fmt.Sprintf("It was expected to have an error message, but it was nil")
// 	}

// }

// // TestValidateInputPointListDifferntDimensions tests the input check of the ValidateInputPointList function.
// // It checks if the function panics when the input has points with different dimensions.
// //
// // Parameters:
// //   - t (*testing.T): A testing object used to run tests and check for failures.
// //
// // Returns:
// //   - None
// func TestValidateInputPointListDifferntDimensions(t *testing.T) {
// 	points := [][]float64{
// 		{1.0, 2.0},
// 		{1.0, 2.0, 3.0},
// 	}

// 	geometry := geom2d.NewEuclid()

// 	errorMsg := geometry.Decimate.ValidateInputPointList(points)

// 	if errorMsg != nil {
// 		fmt.Sprintf("It was expected to have an error message, but it was nil")
// 	}

// }

// // TestValidateDouglasPeuckerOnlyTwoPointsInput tests the DouglasPeucker function.
// // It checks if the function returns the same input when the input has only two points.
// //
// // Parameters:
// //   - t (*testing.T): A testing object used to run tests and check for failures.
// //
// // Returns:
// //   - None
// func TestValidateDouglasPeuckerOnlyTwoPointsInput(t *testing.T) {
// 	points := [][]float64{
// 		{1.0, 2.0},
// 		{1.0, 2.0},
// 	}
// 	expected := [][]float64{
// 		{1.0, 2.0},
// 		{1.0, 2.0},
// 	}

// 	geometry := geom2d.NewEuclid()

// 	points = geometry.Decimate.DouglasPeucker(points, 0.1)

// 	if len(points) != len(expected) {
// 		t.Errorf("Expected %v points, but got %v", len(expected), len(points))
// 	}

// 	errorMsg := ""
// 	for i, point := range points {
// 		for j := range point {
// 			if points[i][j] != expected[i][j] {
// 				errorMsg += fmt.Sprintf("\nAt coordinates (%v,%v)  expected value %v, but got %v", i, j, expected[i][j], points[i][j])
// 			}
// 		}
// 	}
// 	if len(errorMsg) > 0 {
// 		t.Errorf(errorMsg)
// 	}
// }

// // TestDouglasPeuckerAllPointsInsideThreshold tests the DouglasPeucker function.
// // It checks if the function returns the expected output when all points are inside the threshold.
// //
// // Parameters:
// //   - t (*testing.T): A testing object used to run tests and check for failures.
// //
// // Returns:
// //   - None
// func TestDouglasPeuckerAllPointsInsideThreshold(t *testing.T) {
// 	points := [][]float64{
// 		{1.0, 2.0},
// 		{2.0, 2.0},
// 		{3.0, 2.0},
// 		{4.0, 2.0},
// 		{5.0, 2.0},
// 		{6.0, 2.0},
// 		{7.0, 2.0},
// 	}

// 	expected := [][]float64{
// 		{1.0, 2.0},
// 		{7.0, 2.0},
// 	}

// 	geometry := geom2d.NewEuclid()

// 	points = geometry.Decimate.DouglasPeucker(points, 0.1)

// 	errorMsg := ""
// 	for i, point := range points {
// 		for j := range point {
// 			if points[i][j] != expected[i][j] {
// 				errorMsg += fmt.Sprintf("\nAt coordinates (%v,%v)  expected value %v, but got %v", i, j, expected[i][j], points[i][j])
// 			}
// 		}
// 	}
// 	if len(errorMsg) > 0 {
// 		t.Errorf(errorMsg)
// 	}
// }

func TestDouglasPeucker(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{2.0, -2.2},
		{3.0, 2.0},
		{4.0, 2.2},
		{5.0, -2.2},
		{7.0, 2.0},
	}

	expected := [][]float64{
		{1.0, 2.0},
		{2.0, -2.2},
		{3.0, 2.0},
		{4.0, 2.2},
		{5.0, -2.2},
		{7.0, 2.0},
	}

	geometry := geom2d.NewEuclid()

	points = geometry.Decimate.DouglasPeucker(points, .50)

	errorMsg := ""
	for i, point := range points {
		for j := range point {
			if points[i][j] != expected[i][j] {
				errorMsg += fmt.Sprintf("\nAt coordinates (%v,%v)  expected value %v, but got %v", i, j, expected[i][j], points[i][j])
				errorMsg += fmt.Sprintf("points %v", points)
			}
		}
	}
	if len(errorMsg) > 0 {
		t.Errorf(errorMsg)
	}
}
