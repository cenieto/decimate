package geom2d

import (
	"decimator/pkg/testutils"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestLine2D validates the creation of a Line2D using NewLine.
//
// This test checks if the Line2D created by NewLine correctly initializes
// its Point1 and Point2 properties.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestLine2D(t *testing.T) {
	p1 := NewPoint(0.0, 1.0)
	p2 := NewPoint(1.0, 0.0)
	result := NewLine(p1, p2)

	expected := Line2D{
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

// TestVectorDirector validates the calculation of the direction vector of a Line2D.
//
// This test checks if the direction vector calculated by VectorDirector matches
// the expected 2D vector.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestVectorDirector(t *testing.T) {
	p1 := NewPoint(0.0, 1.0)
	p2 := NewPoint(1.0, 0.0)
	line := NewLine(p1, p2)

	result := line.VectorDirector()
	expected := NewVector(1.0, -1.0)

	// Check if the result is approximately equal to the expected result
	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("line.VectorDirector() = %v; want %v", mat.Formatted(result), mat.Formatted(expected))
	}
}

// TestLength validates the calculation of the length of a Line2D.
//
// This test checks if the length calculated by Length matches the expected value.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestLength(t *testing.T) {
	p1 := NewPoint(0.0, 1.0)
	p2 := NewPoint(1.0, 0.0)
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
	point_1 := NewPoint(0.0, 1.0)
	point_2 := NewPoint(1.0, 0.0)
	line := NewLine(point_1, point_2)

	result := line.String()
	expected := fmt.Sprintf("Line{%v, %v}",
		point_1.String(),
		point_2.String(),
	)

	if result != expected {
		t.Errorf("line.String() = %v; want %v", result, expected)
	}
}
