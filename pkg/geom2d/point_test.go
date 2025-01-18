package geom2d

import (
	"decimator/pkg/testutils"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestNewPoint validates the creation of a Point2D using NewPoint.
//
// This test checks if the Point2D created by NewPoint matches the expected
// 2D vector representation.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestNewPoint(t *testing.T) {
	input := []float64{0.0, 1.0}
	result := NewPoint(input[0], input[1])
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
	point := NewPoint(0.0, 1.0)

	result := point.String()
	expected := fmt.Sprintf("Point{%v}", mat.Formatted(point))

	if result != expected {
		t.Errorf("point.String() = %v; want %v", result, expected)
	}
}
