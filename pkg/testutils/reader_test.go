package testutils

import (
	"reflect"
	"testing"
)

// TestNewCSVFloat64Reader tests the NewCSVFloat64Reader function.
// Verifies that the function returns a CSVFloat64Reader with the correct header.
// Arguments: none
// Returns: none
func TestNewCSVFloat64Reader(t *testing.T) {
	fixtureFile := "../../testdata/testutils/reader-example.csv"
	reader, err := NewCSVFloat64Reader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	expectedHeader := []string{"A", "B", "C"}

	if !reflect.DeepEqual(reader.header, expectedHeader) {
		t.Fatalf("Expected header %v, got %v", expectedHeader, reader.header)
	}

	expecteRow := []string{"1.1", "2.2", "3.3"}
	row, err := reader.reader.Read()

	if !reflect.DeepEqual(row, expecteRow) {
		t.Fatalf("Expected row %v, got %v", expecteRow, row)
	}

	expecteRow = []string{"2.1", "2.2", "2.3"}
	row, err = reader.reader.Read()

	if !reflect.DeepEqual(row, expecteRow) {
		t.Fatalf("Expected row %v, got %v", expecteRow, row)
	}

	expected := "../../testdata/testutils/reader-example.csv"
	result := reader.file.Name()

	if result != expected {
		t.Fatalf("Expected file %v, got %v", expected, result)
	}
}

// TestReadLines tests the ReadLines function.
// Verifies that the function reads all lines from the CSV file.
// Arguments: none
// Returns: none
func TestReadLines(t *testing.T) {
	fixtureFile := "../../testdata/testutils/reader-example.csv"
	reader, err := NewCSVFloat64Reader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	expected := [][]float64{
		{1.1, 2.2, 3.3},
		{2.1, 2.2, 2.3},
	}

	lines, _ := reader.ReadLines()

	i := 0
	for line := range lines {
		if !reflect.DeepEqual(line, expected[i]) {
			t.Fatalf("Expected line %v, got %v", expected[i], line)
		}
		i++
	}
}
