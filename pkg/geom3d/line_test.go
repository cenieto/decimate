package geom3d

import (
	"decimator/pkg/testutils"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestLineString(t *testing.T) {
	p1 := NewPoint(0.0, 1., 2.0)
	p2 := NewPoint(1.0, 0.0, 2.0)
	line := NewLine(p1, p2)

	result := line.String()
	expected := fmt.Sprintf("Line3D(%v, %v)",
		p1.String(),
		p2.String(),
	)

	if result != expected {
		t.Errorf("line.String() = %v; want %v", result, expected)
	}
}

// TestLine3D validates the creation of a Line3D using NewLine.
//
// This test checks if the Line3D created by NewLine correctly initializes
// its Point1 and Point2 properties.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestLine3D(t *testing.T) {
	p1 := NewPoint(0.0, 1., 2.0)
	p2 := NewPoint(1.0, 0.0, 2.0)
	result := NewLine(p1, p2)

	expected := Line3D{
		Point1: p1,
		Point2: p2,
	}

	errorMsgs := false
	if result.Point1 != expected.Point1 {
		errorMsgs = true
	}

	if errorMsgs {
		t.Errorf("NewLine(%v, %v) = %v; want %v", p1.String(), p2.String(), result, expected.String())
	}
}

// TestVectorDirector validates the calculation of the direction vector of a Line3D.
//
// This test checks if the direction vector calculated by VectorDirector matches
// the expected 3D vector.
//
// Arguments:
//
//	t (*testing.T): The testing context provided by the Go testing framework.
func TestVectorDirector(t *testing.T) {
	p1 := NewPoint(0.0, 1.0, 2.0)
	p2 := NewPoint(1.0, 0.0, 2.0)
	line := NewLine(p1, p2)

	result := line.VectorDirector()
	expected := NewVector(1.0, -1.0, 0.0)

	// Check if the result is approximately equal to the expected result
	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("line.VectorDirector() = %v; want %v", result.String(), expected.String())
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
	p1 := NewPoint(0.0, 0.0, 0.0)
	p2 := NewPoint(1.0, 1.0, 1.0)
	line := NewLine(p1, p2)

	result := line.Length()
	expected := 1.7320508075688772

	// Check if the result is approximately equal to the expected result
	if result != expected {
		t.Errorf("line.Length() = %v; want %v", result, expected)
	}
}
