package testutils

import (
	"reflect"
	"testing"
)

// TestNewCSVFloat64Reader tests the NewCSVFloat64Reader function.
// This test ensures that the function correctly initializes a CSVFloat64Reader,
// reads the header from the CSV file, and returns the expected header and rows.
// It also verifies that the correct file path is returned.
// Steps:
// 1. Open a test CSV file.
// 2. Check if the returned header matches the expected header.
// 3. Validate the subsequent rows against expected data.
// 4. Confirm the file path matches the expected path.
//
// Arguments: none
// Returns: none
func TestNewCSVFloat64Reader(t *testing.T) {
	fixtureFile := "../../testdata/testutils/reader-example.csv"
	reader, err := NewCSVFloat64Reader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	expectedHeader := []string{"A", "B", "C"}

	// Verify header
	if !reflect.DeepEqual(reader.header, expectedHeader) {
		t.Fatalf("Expected header %v, got %v", expectedHeader, reader.header)
	}

	// Verify rows
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

	// Verify file path
	expected := "../../testdata/testutils/reader-example.csv"
	result := reader.file.Name()

	if result != expected {
		t.Fatalf("Expected file %v, got %v", expected, result)
	}
}

// TestReadLines tests the ReadLines function of the CSVFloat64Reader.
// This test ensures that the ReadLines function correctly parses all lines from
// the CSV file into slices of float64. It validates the parsed values against
// the expected data.
//
// Steps:
// 1. Open a test CSV file.
// 2. Use ReadLines to parse the file's content.
// 3. Verify each parsed line matches the expected data.
//
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

	results := reader.ReadLines()

	// Verify each parsed line
	i := 0
	for result := range results {
		if !reflect.DeepEqual(result.Values, expected[i]) {
			t.Fatalf("Expected result %v, got %v", expected[i], result)
		}
		i++
	}
}

// TestReadLinesErroneousLine tests the ReadLines function with a CSV file
// that contains an erroneous line (e.g., mismatched number of fields).
// This test ensures that the function detects the error and provides
// appropriate error messages.
//
// Steps:
// 1. Open a test CSV file with a line that has an incorrect number of fields.
// 2. Use ReadLines to parse the file's content.
// 3. Verify that the error is detected and matches the expected error message.
//
// Arguments: none
// Returns: none
func TestReadLinesErroneousLine(t *testing.T) {
	fixtureFile := "../../testdata/testutils/reader-example-error-number-of-fields.csv"
	reader, err := NewCSVFloat64Reader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	results := reader.ReadLines()

	var errors []error
	// Collect all errors
	for result := range results {
		if result.Err != nil {
			errors = append(errors, result.Err)
		}
	}

	if len(errors) == 0 {
		t.Fatalf("Expected error, got none")
	}

	// Verify the error message
	expected := string("record on line 3: wrong number of fields")
	if errors[0].Error() != expected {
		t.Fatalf("Expected \"%v\", got \"%v\"", expected, errors[0].Error())
	}
}

// TestReadLinesErrParseValue tests the ReadLines function with a CSV file
// that contains a value that cannot be parsed as a float64.
// This test ensures that the function detects parsing errors and provides
// the appropriate error messages.
//
// Steps:
// 1. Open a test CSV file with a value that cannot be parsed (e.g., non-numeric string).
// 2. Use ReadLines to parse the file's content.
// 3. Verify that the error is detected and matches the expected error message.
//
// Arguments: none
// Returns: none
func TestReadLinesErrParseValue(t *testing.T) {
	fixtureFile := "../../testdata/testutils/reader-example-error-parse-value.csv"
	reader, err := NewCSVFloat64Reader(fixtureFile)
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	results := reader.ReadLines()

	var errors []error
	// Collect all errors
	for result := range results {
		if result.Err != nil {
			errors = append(errors, result.Err)
		}
	}

	if len(errors) == 0 {
		t.Fatalf("Expected error, got none")
	}

	// Verify the error message
	expected := string("error parsing float on line 1: strconv.ParseFloat: parsing \"a\": invalid syntax")
	if errors[0].Error() != expected {
		t.Fatalf("Expected \"%v\", got \"%v\"", expected, errors[0].Error())
	}
}
