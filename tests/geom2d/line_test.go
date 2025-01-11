package test

import (
	"decimator/pkg/geom2d"
	"decimator/tests/testutils"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestLine2D validates the creation of a Line2D using NewLine.
//
// This test checks if the Line2D created by geom2d.NewLine correctly initializes
// its Point1 and Point2 properties.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestLine2D(t *testing.T) {
	p1 := geom2d.NewPoint(0.0, 1.0)
	p2 := geom2d.NewPoint(1.0, 0.0)
	result := geom2d.NewLine(p1, p2)

	expected := geom2d.Line2D{
		Point1: p1,
		Point2: p2,
	}

	errorMsgs := false
	if result.Point1 != expected.Point1 {
		errorMsgs = true
	}

	if errorMsgs {
		t.Errorf("geom2d.NewLine(%v, %v) = %v; want %v", mat.Formatted(p1), mat.Formatted(p2), result, expected)
	}
}

// TestVectorDirector validates the calculation of the direction vector of a Line2D.
//
// This test checks if the direction vector calculated by VectorDirector matches
// the expected 2D vector.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestVectorDirector(t *testing.T) {
	p1 := geom2d.NewPoint(0.0, 1.0)
	p2 := geom2d.NewPoint(1.0, 0.0)
	line := geom2d.NewLine(p1, p2)

	result := line.VectorDirector()
	expected := geom2d.NewVector(1.0, -1.0)

	// Check if the result is approximately equal to the expected result
	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("line.VectorDirector() = %v; want %v", mat.Formatted(result), mat.Formatted(expected))
	}
}
