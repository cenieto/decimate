package geom3d

import (
	"decimator/pkg/testutils"
	"gonum.org/v1/gonum/mat"
	"io"
	"testing"
)

// TestGeom3DInstantiation tests the correct instantiation of a geom3D object.
// Verifies that a newly created instance of geom3D returns a dimension of 3, which is expected for a 3D geometry.
// Arguments: none
// Returns: none
func TestGeom3DInstantiation(t *testing.T) {
	geom3D := NewEuclid() // Create a new geom3D instance

	dimension := geom3D.Dimension() // Get the dimension of the geometry

	expected := 3 // Expected dimension for a 3D system

	// Verify if the returned dimension matches the expected value
	if dimension != expected {
		t.Errorf("geom3D.Dimension() = %v; want %v", dimension, expected)
	}
}

// TestCrossProduct3D tests the cross product operation for 3D vectors.
// Verifies that the cross product of 3D vectors results in a 3D vector.
// Set of inputs and expected results are read from a CSV file.
// This test ensures the correctness of the cross product operation for 3D vectors in a 3D space.
// Arguments: none
// Returns: none
func TestCrossProduct3D(t *testing.T) {

	fixtureFile := "../../testdata/geom3d/cross-product-3d.csv"  // Updated fixture file path
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

		geom := NewEuclid() // Create new geom3D instance

		v1 := NewVector(values[0], values[1], values[2]) // First vector for cross product
		v2 := NewVector(values[3], values[4], values[5]) // Second vector for cross product

		// Compute the cross product of v1 and v2
		result := geom.CrossProduct(v1, v2)

		// Check if the result is approximately equal to the expected result
		if !mat.EqualApprox(expected, result, testutils.TestToleranceRelative) {
			t.Errorf("TestCrossProduct3D(%v, %v) = %v; want %v", mat.Formatted(v1), mat.Formatted(v2), mat.Formatted(result), mat.Formatted(expected))
		}
	}

}

// TestCrossProductNorm tests the cross product norm (magnitude) for 3D vectors.
// Verifies that the cross product of two 3D vectors results in a correct magnitude
// Set of inputs and expected results are read from a CSV file.
// Arguments: none
// Returns: none
func TestCrossProductNorm(t *testing.T) {
	fixtureFile := "../../testdata/geom3d/cross-product-3d.csv"  // Updated fixture file path
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
		expected := values[9] // Expected result for the magnitude of the cross product

		geom := NewEuclid() // Create new geom3D instance

		v1 := NewVector(values[0], values[1], values[2]) // First vector for cross product
		v2 := NewVector(values[3], values[4], values[5]) // Second vector for cross product

		// Compute the magnitude of the cross product of v1 and v2
		result := geom.CrossProductNorm(v1, v2)

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("TestCrossProductNorm(%v, %v) = %v; want %v", mat.Formatted(v1), mat.Formatted(v2), result, expected)
		}
	}
}

// // TestDoubleAreaTriangle tests the calculation of the double area of a triangle formed by a point and a line.
// // Verifies that the double area of a triangle with vertices in three 3D points is computed correctly.
// // Set of inputs and expected results are read from a CSV file.
// // Arguments: none
// // Returns: none
func TestDoubleAreaTriangle(t *testing.T) {
	fixtureFile := "../../testdata/geom3d/point-line-3d.csv"     // Updated fixture file path
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

		geom := NewEuclid()
		point := NewPoint(values[0], values[1], values[2])
		point_origin_line := NewPoint(values[3], values[4], values[5])
		point_end_line := NewPoint(values[6], values[7], values[8])
		line := NewLine(point_origin_line, point_end_line)

		result := geom.DoubleAreaTriangle(point, line)
		expected := values[9]

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("DoubleAreaTriangle(%v, %v) = %v; want %v", mat.Formatted(point), line, result, expected)
		}
	}
}

// // TestDistancePointLine tests the calculation of the shortest distance from a point to a line.
// // Verifies that the distance from one 3D point to the line passing through two 3D points is computed correctly.
// // Set of inputs and expected results are read from a CSV file.
// // Arguments: none
// // Returns: none
func TestDistancePointLine(t *testing.T) {
	fixtureFile := "../../testdata/geom3d/point-line-3d.csv"     // Updated fixture file path
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

		geom := NewEuclid()
		point := NewPoint(values[0], values[1], values[2])
		point_origin_line := NewPoint(values[3], values[4], values[5])
		point_end_line := NewPoint(values[6], values[7], values[8])
		line := NewLine(point_origin_line, point_end_line)

		result := geom.DistancePointLine(point, line)
		expected := values[10]

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("DoubleAreaTriangle(%v, %v) = %v; want %v", mat.Formatted(point), line, result, expected)
		}
	}
}
