// Package decimator provides functionalities to decimate, remove points, from a trajectory
// defined as a set of 4D points.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Point4D is a structure to store spatial/temporal definition of a four dimensional point
type Point4D struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	T float32 `json:"t"`
}

// Segment is a structure to store the definition of a four dimensional segment
type Segment struct {
	Point1 Point4D
	Point2 Point4D
}

// ReadFile opens and reads a file and returns its contents
// 
// Parameters:
//	- inputFile: path to the input file
//
// Returns:
//	- the content of the input file as an array of bytes
func ReadFile(inputFile string) ([]byte, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error while opening file: %v", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	return content, nil
}

// ReadPoints opens and reads a file and returns its contents as an array of Point4D
//
// Parameters:
//	- inputFile: path to the input file
//
// Returns:
//	- the content of the input file as an array of Point4D
func readPoints(filename string) []Point4D {

	bytes, err := ReadFile(filename)
	var points []Point4D
	err = json.Unmarshal(bytes, &points)
	if err != nil {
		log.Fatalf("Error while unmarshalling JSON: %v", err)
	}
	return points
}

func main() {
	inputFile := "./fixtures/input-4d.json"

	points := readPoints(inputFile)

	fmt.Println(points)

}
