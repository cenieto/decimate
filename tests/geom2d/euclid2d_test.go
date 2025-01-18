package test

import (
	"decimator/pkg/geom2d"
	"decimator/tests/testutils"
	"gonum.org/v1/gonum/mat"
	"io"
	"math"
	"testing"
)

// TestGeom2dInstantiation tests the correct instantiation of a geom2d object.
// Verifies that a newly created instance of geom2d returns a dimension of 2, which is expected for a 2D geometry.
// Arguments: none
// Returns: none
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
// Verifies that the cross product of 2D vectors results in a 3D vector.
// Set of inputs and expected results are read from a CSV file.
// This test ensures the correctness of the cross product operation for 2D vectors in a 3D space.
// Arguments: none
// Returns: none
func TestCrossProduct2D(t *testing.T) {

	fixtureFile := "../../testdata/geom2d/cross-product.csv"  // Updated fixture file path
	reader, err := testutils.NewCSVFloat64Reader(fixtureFile) // Get the CSV float64 reader
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Error while reading line %v\n", err)
		}

		expected := mat.NewVecDense(3, []float64{values[6], values[7], values[8]}) // Expected result for the cross product

		geom := geom2d.NewEuclid() // Create new geom2d instance

		v1 := geom2d.NewVector(values[0], values[1]) // First vector for cross product
		v2 := geom2d.NewVector(values[3], values[4]) // Second vector for cross product

		// Compute the cross product of v1 and v2
		result := geom.CrossProduct(v1, v2)

		// Check if the result is approximately equal to the expected result
		if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
			t.Errorf("TestCrossProduct2D(%v, %v) = %v; want %v", mat.Formatted(v1), mat.Formatted(v2), mat.Formatted(result), mat.Formatted(expected))
		}
	}

}

// TestCrossProductNorm tests the cross product norm (magnitude) for 2D vectors.
// Verifies that the cross product of two 2D vectors results in a correct magnitude
// Set of inputs and expected results are read from a CSV file.
// Arguments: none
// Returns: none
func TestCrossProductNorm(t *testing.T) {
	fixtureFile := "../../testdata/geom2d/cross-product.csv"  // Updated fixture file path
	reader, err := testutils.NewCSVFloat64Reader(fixtureFile) // Get the CSV float64 reader
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Error while reading line %v\n", err)
		}
		expected := math.Abs(values[8]) // Expected result for the magnitude of the cross product

		geom := geom2d.NewEuclid() // Create new geom2d instance

		v1 := geom2d.NewVector(values[0], values[1]) // First vector for cross product
		v2 := geom2d.NewVector(values[3], values[4]) // Second vector for cross product

		// Compute the magnitude of the cross product of v1 and v2
		result := geom.CrossProductNorm(v1, v2)

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("TestCrossProductNorm(%v, %v) = %v; want %v", mat.Formatted(v1), mat.Formatted(v2), result, expected)
		}
	}
}

// TestDoubleAreaTriangle tests the calculation of the double area of a triangle formed by a point and a line.
// Verifies that the double area of a triangle with vertices in three 2D points is computed correctly.
// Set of inputs and expected results are read from a CSV file.
// Arguments: none
// Returns: none
func TestDoubleAreaTriangle(t *testing.T) {
	fixtureFile := "../../testdata/geom2d/point-line.csv"     // Updated fixture file path
	reader, err := testutils.NewCSVFloat64Reader(fixtureFile) // Get the CSV float64 reader
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Error while reading line %v\n", err)
		}

		geom := geom2d.NewEuclid()
		point := geom2d.NewPoint(values[0], values[1])
		point_origin_line := geom2d.NewPoint(values[2], values[3])
		point_end_line := geom2d.NewPoint(values[4], values[5])
		line := geom2d.NewLine(point_origin_line, point_end_line)

		result := geom.DoubleAreaTriangle(point, line)
		expected := values[6]

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("DoubleAreaTriangle(%v, %v) = %v; want %v", point, line, result, expected)
		}
	}
}

// TestDistancePointLine tests the calculation of the shortest distance from a point to a line.
// Verifies that the distance from one 2D point to the line passing through two 2D points is computed correctly.
// Set of inputs and expected results are read from a CSV file.
// Arguments: none
// Returns: none
func TestDistancePointLine(t *testing.T) {
	fixtureFile := "../../testdata/geom2d/point-line.csv"     // Updated fixture file path
	reader, err := testutils.NewCSVFloat64Reader(fixtureFile) // Get the CSV float64 reader
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Error while reading line %v\n", err)
		}

		geom := geom2d.NewEuclid()
		point := geom2d.NewPoint(values[0], values[1])
		point_origin_line := geom2d.NewPoint(values[2], values[3])
		point_end_line := geom2d.NewPoint(values[4], values[5])
		line := geom2d.NewLine(point_origin_line, point_end_line)

		result := geom.DistancePointLine(point, line)
		expected := values[7]

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("DoubleAreaTriangle(%v, %v) = %v; want %v", point, line, result, expected)
		}
	}
}
