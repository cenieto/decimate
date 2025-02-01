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
	"decimator/pkg/testutils"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestNewVector checks the creation of a new Vector object.
func TestNewVector(t *testing.T) {
	input := []float64{0.0, 1.0}
	result := NewVector(input)
	expected := mat.NewVecDense(2, input)

	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("NewVector(%v) = %v; want %v", input, mat.Formatted(result), mat.Formatted(expected))
	}
}

// TestNewVectorTwoPointsPanic checks the behavior of NewVectorTwoPoints when dimensions differ.
func TestNewVectorTwoPointsPanic(t *testing.T) {
	tests := []struct {
		name            string
		point1          *Point
		point2          *Point
		expectedPanic   bool
		expectedMessage string
	}{
		{
			name:            "SameDimensions",
			point1:          NewPoint([]float64{1.0, 2.0}),
			point2:          NewPoint([]float64{3.0, 4.0}),
			expectedPanic:   false,
			expectedMessage: "",
		},
		{
			name:            "DifferentDimensions",
			point1:          NewPoint([]float64{1.0, 2.0}),
			point2:          NewPoint([]float64{3.0, 4.0, 5.0}),
			expectedPanic:   true,
			expectedMessage: "Points must have the same dimension: 2 != 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectedPanic {
						t.Errorf("Unexpected panic: %v", r)
					} else {
						if r != tt.expectedMessage {
							t.Errorf("Expected panic message: %q, got: %q", tt.expectedMessage, r)
						}
					}
				} else if tt.expectedPanic {
					t.Errorf("Expected panic but none occurred")
				}
			}()

			_ = NewVectorTwoPoints(tt.point1, tt.point2)
		})
	}
}

// TestVectorString checks the string representation of a Vector.
func TestVectorString(t *testing.T) {
	vector := NewVector([]float64{0.0, 1.0})

	result := vector.String()
	expected := fmt.Sprintf("%v", mat.Formatted(vector.VecDense))

	if result != expected {
		t.Errorf("vector.String() = %v; want %v", result, expected)
	}
}

// TestVectorLength checks the computation of the Euclidean length of a Vector.
func TestVectorLength(t *testing.T) {
	vector := NewVector([]float64{1.0, 1.0})

	result := vector.Length()
	expected := 1.4142135623730951

	if result != expected {
		t.Errorf("vector.Length() = %v; want %v", result, expected)
	}
}

// TestVectorDimension checks the dimension of a Vector.
func TestVectorDimension(t *testing.T) {
	vector := NewVector([]float64{1.0, 1.0})

	result := vector.Dimension()
	expected := 2

	if result != expected {
		t.Errorf("vector.Dimension() = %v; want %v", result, expected)
	}
}
