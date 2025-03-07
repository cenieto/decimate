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
package primitives

import (
	"fmt"
	"github.com/cenieto/decimate/pkg/testutils"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestLine validates the creation of a Line using NewLine.
//
// This test checks if the Line created by NewLine correctly initializes
// its Point1 and Point2 properties.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestLine(t *testing.T) {
	p1 := NewPoint([]float64{0.0, 1.0})
	p2 := NewPoint([]float64{1.0, 0.0})
	result := NewLine(p1, p2)

	expected := Line{
		Point1: p1,
		Point2: p2,
	}

	errorMsgs := false
	if result.Point1 != expected.Point1 {
		errorMsgs = true
	}

	if errorMsgs {
		t.Errorf("NewLine(%v, %v) = %v; want %v", mat.Formatted(p1), mat.Formatted(p2), result, expected)
	}
}

// TestVectorDirector validates the calculation of the direction vector of a Line.
//
// This test checks if the direction vector calculated by VectorDirector matches
// the expected 2D vector.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestVectorDirector(t *testing.T) {
	p1 := NewPoint([]float64{0.0, 1.0})
	p2 := NewPoint([]float64{1.0, 0.0})
	line := NewLine(p1, p2)

	result := line.VectorDirector()
	expected := NewVector([]float64{1.0, -1.0})

	// Check if the result is approximately equal to the expected result
	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("line.VectorDirector() = %v; want %v", mat.Formatted(result), mat.Formatted(expected))
	}
}

// TestLength validates the calculation of the length of a Line.
//
// This test checks if the length calculated by Length matches the expected value.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestLength(t *testing.T) {
	p1 := NewPoint([]float64{0.0, 0.0})
	p2 := NewPoint([]float64{1.0, 1.0})
	line := NewLine(p1, p2)

	result := line.Length()
	expected := 1.4142135623730951

	// Check if the result is approximately equal to the expected result
	if result != expected {
		t.Errorf("line.Length() = %v; want %v", result, expected)
	}
}

// TestString validates the string representation of a Line.
//
// This test checks if the string representation of a Line object matches the expected format.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestLineString(t *testing.T) {
	p1 := NewPoint([]float64{0.0, 0.0})
	p2 := NewPoint([]float64{1.0, 1.0})
	line := NewLine(p1, p2)

	result := line.String()
	expected := fmt.Sprintf("Line{%v, %v}",
		p1.String(),
		p2.String(),
	)

	if result != expected {
		t.Errorf("line.String() = %v; want %v", result, expected)
	}
}

// TestLineDimension validates the dimension of a Point2D.
//
// This test checks if the dimension of a Point2D matches the expected value.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestLineDimension(t *testing.T) {
	p1 := NewPoint([]float64{0.0, 0.0})
	p2 := NewPoint([]float64{1.0, 1.0})
	line := NewLine(p1, p2)

	result := line.Dimension()
	expected := 2

	if result != expected {
		t.Errorf("line.Dimension() = %v; want %v", result, expected)
	}
}
