package geom3d

import (
	"decimator/pkg/testutils"
	"gonum.org/v1/gonum/mat"
	"testing"
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

	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("NewPoint(%v) = %v; want %v", input, mat.Formatted(result), mat.Formatted(expected))
	}
}
