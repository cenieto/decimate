// Copyright 2025 César Nieto Sánchez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package primitives

import (
	"decimate/pkg/testutils"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestNewPoint validates the creation of a Point using NewPoint.
//
// This test checks if the Point created by NewPoint matches the expected
// vector representation.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestNewPoint(t *testing.T) {
	input := []float64{0.0, 1.0}
	result := NewPoint(input)
	expected := mat.NewVecDense(2, input)

	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("NewPoint(%v) = %v; want %v", input, mat.Formatted(result), mat.Formatted(expected))
	}
}

// TestString validates the string representation of a Point.
//
// This test checks if the string representation of a Point matches the
// expected format.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestPointString(t *testing.T) {
	input := []float64{0.0, 1.0}
	point := NewPoint(input)

	result := point.String()
	expected := fmt.Sprintf("Point{%v}", mat.Formatted(point))

	if result != expected {
		t.Errorf("point.String() = %v; want %v", result, expected)
	}
}

// TestPointDimension validates the dimension of a Point2D.
//
// This test checks if the dimension of a Point2D matches the expected value.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestPointDimension(t *testing.T) {
	input := []float64{0.0, 1.0}
	point := NewPoint(input)

	result := point.Dimension()
	expected := 2

	if result != expected {
		t.Errorf("point.Dimension() = %v; want %v", result, expected)
	}
}
