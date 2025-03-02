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
	"github.com/cenieto/decimate/pkg/geom3d"
	"github.com/cenieto/decimate/pkg/testutils"
	"testing"
)

// TestValidateInputPointListOnePoint tests the input check of the ValidateInputPointList function.
// It checks if the function panics when the input has only one element.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestValidateInputPointListOnePoint(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
	}

	geometry := geom2d.NewEuclid()

	errorMsg := geometry.Decimate.ValidateInputPointList(points)
	if errorMsg != nil {
		t.Errorf("It was expected to have an error message, but it was nil")
	}

}

// TestValidateInputPointListDifferntDimensions tests the input check of the ValidateInputPointList function.
// It checks if the function panics when the input has points with different dimensions.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestValidateInputPointListDifferntDimensions(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{1.0, 2.0, 3.0},
	}

	geometry := geom2d.NewEuclid()

	errorMsg := geometry.Decimate.ValidateInputPointList(points)

	if errorMsg != nil {
		t.Errorf("It was expected to have an error message, but it was nil")
	}

}

// TestValidateDouglasPeuckerOnlyTwoPointsInput tests the DouglasPeucker function.
// It checks if the function returns the same input when the input has only two points.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestValidateDouglasPeuckerOnlyTwoPointsInput(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{1.0, 2.0},
	}
	expected := [][]float64{
		{1.0, 2.0},
		{1.0, 2.0},
	}

	geometry := geom2d.NewEuclid()

	points = geometry.Decimate.DouglasPeucker(points, 0.1)
	result, error := testutils.CompareSlices(points, expected)
	if !result {
		t.Errorf("The test failed, %v", error)
	}
}

// TestDouglasPeuckerAllPointsInsideThreshold tests the DouglasPeucker function.
// It checks if the function returns the expected output when all points are inside the threshold.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestDouglasPeuckerAllPointsInsideThreshold(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{2.0, 2.0},
		{3.0, 2.0},
		{4.0, 2.0},
		{5.0, 2.0},
		{6.0, 2.0},
		{7.0, 2.0},
	}

	expected := [][]float64{
		{1.0, 2.0},
		{7.0, 2.0},
	}

	geometry := geom2d.NewEuclid()

	points = geometry.Decimate.DouglasPeucker(points, 0.1)

	result, error := testutils.CompareSlices(points, expected)
	if !result {
		t.Errorf("The test failed, %v", error)
	}

}

// TestDouglasPeuckerSingleLineFixedOffset tests the DouglasPeucker function.
// It checks if the function returns the expected output when the input is a single line with fixed offset.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestDouglasPeuckerSingleLineFixedOffset(t *testing.T) {
	fixtureFile := "../../../testdata/douglas_peucker/single_line_2d_fixed_offset.json"
	data, err := testutils.JSONTestDataReader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening JSON file: %v", err)
	}

	geometry := geom2d.NewEuclid()

	input := data.Input

	for _, test := range data.Expected {
		inputCopy := make([][]float64, len(input))
		copy(inputCopy, input)
		points := geometry.Decimate.DouglasPeucker(inputCopy, test.Epsilon)
		expected := test.Data
		result, error := testutils.CompareSlices(points, expected)
		if !result {
			t.Errorf("The test failed, %v, expected: %v, result: %v", error, expected, points)
		}

	}
}

// TestDouglasPeuckerPolyineFixedOffset tests the DouglasPeucker function.
// It checks if the function returns the expected output when the input is a polyline with fixed offset.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestDouglasPeuckerPolyineFixedOffset(t *testing.T) {
	fixtureFile := "../../../testdata/douglas_peucker/polyline_2d_fixed_offset.json"
	data, err := testutils.JSONTestDataReader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening JSON file: %v", err)
	}

	geometry := geom2d.NewEuclid()

	input := data.Input

	for _, test := range data.Expected {
		inputCopy := make([][]float64, len(input))
		copy(inputCopy, input)
		points := geometry.Decimate.DouglasPeucker(inputCopy, test.Epsilon)
		expected := test.Data
		result, error := testutils.CompareSlices(points, expected)
		if !result {
			t.Errorf("The test failed, %v, expected: %v, result: %v", error, expected, points)
		}

	}
}

// TestDouglasPeuckerSingleLineRandomOffset tests the DouglasPeucker function.
// It checks if the function returns the expected output when the input is a single line with random offset.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestDouglasPeuckerSingleLineRandomOffset(t *testing.T) {
	fixtureFile := "../../../testdata/douglas_peucker/single_line_2d_noise.json"
	data, err := testutils.JSONTestDataReader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening JSON file: %v", err)
	}

	geometry := geom2d.NewEuclid()

	input := data.Input

	for _, test := range data.Expected {
		inputCopy := make([][]float64, len(input))
		copy(inputCopy, input)
		points := geometry.Decimate.DouglasPeucker(inputCopy, test.Epsilon)
		expected := test.Data
		result, error := testutils.CompareSlices(points, expected)
		if !result {
			t.Errorf("The test failed with epsilon %v, %v, expected: %v\n, result: %v", test.Epsilon, error, expected, points)
		}

	}
}

// TestDouglasPeuckerPolyineRandomOffset tests the DouglasPeucker function.
// It checks if the function returns the expected output when the input is a polyline with random offset.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestDouglasPeucker2DPolyineRandomOffset(t *testing.T) {
	fixtureFile := "../../../testdata/douglas_peucker/polyline_2d_noise.json"
	data, err := testutils.JSONTestDataReader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening JSON file: %v", err)
	}

	geometry := geom2d.NewEuclid()

	input := data.Input

	for _, test := range data.Expected {
		inputCopy := make([][]float64, len(input))
		copy(inputCopy, input)
		points := geometry.Decimate.DouglasPeucker(inputCopy, test.Epsilon)
		expected := test.Data
		result, error := testutils.CompareSlices(points, expected)
		if !result {
			t.Errorf("The test failed with epsilon %v, %v, expected: %v\n, result: %v", test.Epsilon, error, expected, points)
		}

	}
}

func TestDouglasPeucker3DPolyineRandomOffset(t *testing.T) {
	fixtureFile := "../../../testdata/douglas_peucker/polyline_3d_noise.json"
	data, err := testutils.JSONTestDataReader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening JSON file: %v", err)
	}

	geometry := geom3d.NewEuclid()

	input := data.Input

	for _, test := range data.Expected {
		inputCopy := make([][]float64, len(input))
		copy(inputCopy, input)
		points := geometry.Decimate.DouglasPeucker(inputCopy, test.Epsilon)
		expected := test.Data
		result, error := testutils.CompareSlices(points, expected)
		if !result {
			t.Errorf("The test failed with epsilon %v, %v, expected: %v\n, result: %v", test.Epsilon, error, expected, points)
		}

	}
}
