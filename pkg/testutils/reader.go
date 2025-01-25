package testutils

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
	"io"
)

// LineData represents a single line in the CSV file, including its line number and content.
type LineData struct {
	LineNumber int       // Line number in the CSV file (starting from 1)
	Content    []float64 // Parsed content of the line as a slice of float64
}

// CSVFloat64Reader reads a CSV file and returns a slice of float64 values for each line.
type CSVFloat64Reader struct {
	reader *csv.Reader
	file   *os.File
	header []string
}

// NewCSVFloat64Reader creates a new CSVFloat64Reader.
// Arguments:
//   - filePath (string): The path to the CSV file.
//
// Returns:
//   - (*CSVFloat64Reader, error): A pointer to the CSVFloat64Reader and an error if any.
func NewCSVFloat64Reader(filePath string) (*CSVFloat64Reader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	return &CSVFloat64Reader{reader: reader, file: file, header: header}, nil
}

// ReadLines reads the CSV file and returns a channel to iterate over the rows.
//
// Returns:
//
//	(<-chan []float64, <-chan error): Two channels, one for the parsed rows as slices of float64,
//	and one for potential errors during reading.
func (r *CSVFloat64Reader) ReadLines() (<-chan []float64, <-chan error) {
	linesChannel := make(chan []float64)
	errorsChannel := make(chan error)

	go func() {
		defer close(linesChannel)
		defer close(errorsChannel)

		lineNumber := 1
		for {
			line, err := r.reader.Read()

			// Handle end of file
			if err == io.EOF {
				break
			}

			if err != nil {
				if errors.Is(err, csv.ErrFieldCount) {
					errorsChannel <- err
					return
				}
				break
			}

			if len(line) == 0 {
				continue
			}
			lineNumber++
			// Check if the line has the same number of elements as the header
			if len(line) != len(r.header) {
				errorsChannel <- errors.New(
					"line " + strconv.Itoa(lineNumber) + " does not match the header column count",
				)
				return
			}
			// TODO add test for this

			var values []float64
			for _, v := range line {
				v = strings.TrimSpace(v)
				value, convErr := strconv.ParseFloat(v, 64)
				if convErr != nil {
					errorsChannel <- errors.New(
						"error parsing float on line " + strconv.Itoa(lineNumber) + ": " + convErr.Error(),
					)
					return
				}
				values = append(values, value)
			}
			linesChannel <- values

		}
	}()

	return linesChannel, errorsChannel
}
