package geom3d

import (
	"decimator/pkg/testutils"
	"gonum.org/v1/gonum/mat"
	"testing"

	"fmt"
)

// TestNewPoint validates the creation of a Point3D using NewPoint.
//
// This test checks if the Point3D created by NewPoint matches the expected
// 3D vector representation.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestNewPoint(t *testing.T) {
	input := []float64{0.0, 1.0, 2.0}
	result := NewPoint(input[0], input[1], input[2])
	expected := mat.NewVecDense(3, input)

	if !mat.EqualApprox(expected, result.VecDense, testutils.TestToleranceRelative) {
		t.Errorf("NewPoint(%v) = %v; want %v", input, mat.Formatted(result.VecDense), mat.Formatted(expected))
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
	point := NewPoint(0.0, 1.0, 2.0)

	result := point.String()
	expected := fmt.Sprintf("Point3D{%v}", mat.Formatted(point))

	if result != expected {
		t.Errorf("point.String() = %v; want %v", result, expected)
	}
}
