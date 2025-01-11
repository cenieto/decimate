package test

import (
	"decimator/pkg/geom2d"
	"gonum.org/v1/gonum/mat"
	"testing"
	"decimator/tests/testutils"
)

// TestGeom2dInstantiation tests the correct instantiation of a geom2d object.
// Verifies that a newly created instance of geom2d returns a dimension of 2, which is expected for a 2D geometry.
func TestGeom2dInstantiation(t *testing.T) {
	geom2d := geom2d.NewEuclid() // Create a new geom2d instance

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

	geom := geom2d.NewEuclid() // Create new geom2d instance

	v1 := geom2d.NewVector(1, 0) // First vector for cross product
	v2 := geom2d.NewVector(0, 1) // Second vector for cross product

	// Compute the cross product of v1 and v2
	result := geom.CrossProduct(v1, v2)

	// Check if the result is approximately equal to the expected result
	if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
		t.Errorf("TestCrossProduct2D(%v, %v) = %v; want %v", mat.Formatted(v1), mat.Formatted(v2), mat.Formatted(result), mat.Formatted(expected))
	}
}

// TestCrossProductNorm tests the cross product norm (magnitude) for 2D vectors.
// Verifies that the cross product of (1, 0) and (0, 1) results in a magnitude of 1.0.
func TestCrossProductNorm(t *testing.T) {
	expected := 1.0 // Expected result for the magnitude of the cross product

	geom := geom2d.NewEuclid() // Create new geom2d instance

	v1 := geom2d.NewVector(1, 0) // First vector for cross product
	v2 := geom2d.NewVector(0, 1) // Second vector for cross product

	// Compute the magnitude of the cross product of v1 and v2
	result := geom.CrossProductNorm(v1, v2)

	// Check if the result is approximately equal to the expected magnitude
	if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
		t.Errorf("TestCrossProduct2D(%v, %v) = %v; want %v", mat.Formatted(v1), mat.Formatted(v2), result, expected)
	}
}

// TestDoubleAreaTriangle tests the calculation of the double area of a triangle formed by a point and a line.
// Verifies that the double area of a triangle with vertices (2, 0), (0, 0), and (0, 2) is 4.0.
func TestDoubleAreaTriangle(t *testing.T) {
	geom := geom2d.NewEuclid()
	point := geom2d.NewPoint(2, 0)
	point_origin_line := geom2d.NewPoint(0, 0)
	point_end_line := geom2d.NewPoint(0, 2)
	line := geom2d.NewLine(point_origin_line, point_end_line)

	result := geom.DoubleAreaTriangle(point, line)
	expected := 4.0

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("geom2d.DoubleAreaTriangle(%v,line(%v, %v)) = %v; want %v", mat.Formatted(point), mat.Formatted(line.Point1), mat.Formatted(line.Point2), result, expected)
	}
}

// TestDistancePointLine tests the calculation of the shortest distance from a point to a line.
// Verifies that the distance from the point (2, 0) to the line passing through (0, 0) and (0, 2) is 2.0.
func TestDistancePointLine(t *testing.T) {
	geom := geom2d.NewEuclid()
	point := geom2d.NewPoint(2, 0)
	point_origin_line := geom2d.NewPoint(0, 0)
	point_end_line := geom2d.NewPoint(0, 2)
	line := geom2d.NewLine(point_origin_line, point_end_line)

	result := geom.DistancePointLine(point, line)
	expected := 2.0

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("geom2d.DistancePointLine(%v,line(%v, %v)) = %v; want %v", mat.Formatted(point), mat.Formatted(line.Point1), mat.Formatted(line.Point2), result, expected)
	}
}
