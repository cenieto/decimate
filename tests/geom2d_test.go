package test

import (
	"decimator/pkg/geom2d"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestGeom2dInstantiation tests the correct instantiation of a geom2d object.
// Verifies that a newly created instance of geom2d returns a dimension of 2, which is expected for a 2D geometry.
func TestGeom2dInstantiation(t *testing.T) {
	geom2d := geom2d.NewGeometry() // Create a new geom2d instance

	dimension := geom2d.Dimension() // Get the dimension of the geometry

	expected := 2 // Expected dimension for a 2D system

	// Verify if the returned dimension matches the expected value
	if dimension != expected {
		t.Errorf("geom2d.Dimension() = %v; want %v", dimension, expected)
	}
}

// TestCrossProduct2D tests the cross product operation for 2D vectors.
// Verifies that the cross product of two 2D vectors (1, 0) and (0, 1) results in the 3D vector (0, 0, 1).
// This test ensures the correctness of the cross product operation for 2D vectors in a 3D space.
func TestCrossProduct2D(t *testing.T) {
	expected := mat.NewVecDense(3, []float64{0, 0, 1}) // Expected result for the cross product

	geom := geom2d.NewGeometry() // Create new geom2d instance

	v1 := geom2d.NewVector(1, 0) // First vector for cross product
	v2 := geom2d.NewVector(0, 1) // Second vector for cross product

	// Compute the cross product of v1 and v2
	result := geom.CrossProduct(v1, v2)

	// Check if the result is approximately equal to the expected result
	if !mat.EqualApprox(expected, result, testToleranceRelative) {
		t.Errorf("TestCrossProduct2D(%v, %v) = %v; want %v", mat.Formatted(v1), mat.Formatted(v2), mat.Formatted(result), mat.Formatted(expected))
	}
}

// TestCrossProductNorm tests the cross product norm (magnitude) for 2D vectors.
// Verifies that the cross product of (1, 0) and (0, 1) results in a magnitude of 1.0.
func TestCrossProductNorm(t *testing.T) {
	expected := 1.0 // Expected result for the magnitude of the cross product

	geom := geom2d.NewGeometry() // Create new geom2d instance

	v1 := geom2d.NewVector(1, 0) // First vector for cross product
	v2 := geom2d.NewVector(0, 1) // Second vector for cross product

	// Compute the magnitude of the cross product of v1 and v2
	result := geom.CrossProductNorm(v1, v2)

	// Check if the result is approximately equal to the expected magnitude
	if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testToleranceRelative) {
		t.Errorf("TestCrossProduct2D(%v, %v) = %v; want %v", mat.Formatted(v1), mat.Formatted(v2), result, expected)
	}
}

func TestDistancePointLine(t *testing.T) {
	geom := geom2d.NewGeometry() // Create new geom2d instance
	p1 := geom2d.NewPoint(1, 0)  // First vector for cross product
	p2 := geom2d.NewPoint(0, 0)  // First vector for cross product
	p3 := geom2d.NewPoint(0, 1)  // First vector for cross product
	line := geom2d.NewLine(p2, p3)

	result := geom.DistancePointLine(p1, line)
	expected := 1.0

	if result != expected {
		t.Errorf("geom2d.DistancePointLine")
	}
}
