package primitives

import (
	"decimator/pkg/testutils"
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
