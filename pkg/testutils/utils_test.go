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
	"testing"
)

// TestCompareSlicesDifferentLength tests the Compare2DSlices function.
// It checks if the function returns an error when the input slices have different lengths.
//
// Arguments: none
// Returns: none
func TestCompareSlicesEqual(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}
	expected := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}

	result, _ := CompareDSlices(points, expected)
	if !result {
		t.Fatalf("Error while comparing slices")
	}
}

// TestCompareSlicesDifferentLength tests the Compare2DSlices function.
// It checks if the function returns an error when the input slices have different lengths.
//
// Arguments: none
// Returns: none
func TestCompareSlicesDifferentLength(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
		{3.0, 4.0},
	}
	expected := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}

	result, error := Compare2DSlices(points, expected)
	if result {
		t.Fatalf("Error while comparing slices, returned %v and must be false", result)
	}

	if error == nil {
		t.Fatalf("Error while comparing slices, returned \"%v\" and must be nil", error)
	}
}

// TestCompareSlicesDifferentContent tests the Compare2DSlices function.
// It checks if the function returns an error when the input slices have different content.
//
// Arguments: none
// Returns: none
func TestCompareSlicesDifferentContent(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{3.0, 5.0},
	}
	expected := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}

	result, error := Compare2DSlices(points, expected)
	if result {
		t.Fatalf("Error while comparing slices, returned %v and must be false", result)
	}

	if error == nil {
		t.Fatalf("Error while comparing slices, returned \"%v\" and must be nil", error)
	}
}
