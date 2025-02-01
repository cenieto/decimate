// Copyright 2025 César Nieto Sánchez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package testutils

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// LineData represents a single line in the CSV file, including its line number and content.
// Fields:
//   - LineNumber (int): Line number in the CSV file (starting from 1).
//   - Values ([]float64): Parsed content of the line as a slice of float64.
//   - Err (error): Error encountered while parsing the line.
type LineData struct {
	LineNumber int       // Line number in the CSV file (starting from 1)
	Values     []float64 // Parsed content of the line as a slice of float64
	Err        error     // Error encountered while parsing the line
}

// CSVFloat64Reader reads a CSV file and returns a slice of float64 values for each line.
// Fields:
//   - reader (*csv.Reader): The CSV reader for the file.
//   - file (*os.File): The file being read.
//   - header ([]string): The header row of the CSV file.
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
// Returns:
//   - (<-chan LineData): A channel that emits LineData for each row in the file.
func (r *CSVFloat64Reader) ReadLines() <-chan LineData {
	result := make(chan LineData)

	go func() {
		defer close(result)

		lineNumber := 1
		for {
			line, err := r.reader.Read()

			// Handle end of file
			if err == io.EOF {
				break
			}

			if err != nil {
				if errors.Is(err, csv.ErrFieldCount) {
					result <- LineData{
						LineNumber: lineNumber,
						Values:     []float64{0.0},
						Err:        err,
					}
					return
				}
				break
			}

			if len(line) == 0 {
				continue
			}

			// Check if the line has the same number of elements as the header
			if len(line) != len(r.header) {
				result <- LineData{
					LineNumber: lineNumber,
					Values:     nil,
					Err:        errors.New("line " + strconv.Itoa(lineNumber) + " does not match the header column count"),
				}
				return
			}

			var values []float64
			for _, v := range line {
				v = strings.TrimSpace(v)
				value, convErr := strconv.ParseFloat(v, 64)
				if convErr != nil {
					result <- LineData{
						LineNumber: lineNumber,
						Values:     nil,
						Err:        errors.New("error parsing float on line " + strconv.Itoa(lineNumber) + ": " + convErr.Error()),
					}
					return
				}
				values = append(values, value)
			}
			lineNumber++
			result <- LineData{
				LineNumber: lineNumber,
				Values:     values,
				Err:        nil,
			}

		}
	}()

	return result
}
