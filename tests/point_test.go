package test

import (
	"decimator/pkg/geom2d"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestNewPoint validates the creation of a Point2D using NewPoint.
//
// This test checks if the Point2D created by geom2d.NewPoint matches the expected
// 2D vector representation.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestNewPoint(t *testing.T) {
	input := []float64{0.0, 1.0}
	result := geom2d.NewPoint(input[0], input[1])
	expected := mat.NewVecDense(2, input)

	if !mat.EqualApprox(expected, result, testToleranceRelative) {
		t.Errorf("geom2d.NewPoint(%v) = %v; want %v", input, mat.Formatted(result), mat.Formatted(expected))
	}
}
