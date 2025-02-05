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
// This project uses the Gonum package (https://gonum.org), which is licensed under
// the BSD 3-Clause License:

// Copyright 2013 The Gonum Authors. All rights reserved.

// Redistribution and use in source and binary forms, with or without modification, 
// are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this 
//    list of conditions and the following disclaimer.

// 2. Redistributions in binary form must reproduce the above copyright notice, 
//    this list of conditions and the following disclaimer in the documentation 
//    and/or other materials provided with the distribution.

// 3. Neither the name of the copyright holder nor the names of its contributors 
//    may be used to endorse or promote products derived from this software 
//    without specific prior written permission.
// This project uses the Gonum package (https://gonum.org), which is licensed under
// the BSD 3-Clause License:

// Copyright 2013 The Gonum Authors. All rights reserved.

// Redistribution and use in source and binary forms, with or without modification, 
// are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this 
//    list of conditions and the following disclaimer.

// 2. Redistributions in binary form must reproduce the above copyright notice, 
//    this list of conditions and the following disclaimer in the documentation 
//    and/or other materials provided with the distribution.

// 3. Neither the name of the copyright holder nor the names of its contributors 
//    may be used to endorse or promote products derived from this software 
//    without specific prior written permission.
This project uses the Gonum package (https://gonum.org), which is licensed under
the BSD 3-Clause License:

Copyright 2013 The Gonum Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, 
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this 
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, 
   this list of conditions and the following disclaimer in the documentation 
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors 
   may be used to endorse or promote products derived from this software 
   without specific prior written permission.
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
