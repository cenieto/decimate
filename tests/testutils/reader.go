package testutils

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

// CSVFloat64Reader reads a CSV file and returns a slice of float64 values for each line.
type CSVFloat64Reader struct {
	reader *csv.Reader
	file   *os.File
	header []string
}

// NewCSVFloat64Reader creates a new CSVFloat64Reader.
// Arguments: 
//   - filePath (string): The path to the CSV file.
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

// Read reads the next line from the CSV file and returns a slice of float64 values.
// Arguments: none
// Returns: 
//   - ([]float64, error): A slice of float64 values and an error if any.
func (r *CSVFloat64Reader) Read() ([]float64, error) {
	line, err := r.reader.Read()
	if err != nil {
		if err == io.EOF {
			r.file.Close()
			return nil, err
		}
		return nil, err
	}

	var values []float64
	for _, v := range line {
		v = strings.TrimSpace(v)
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}
