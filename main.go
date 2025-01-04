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

// LineSegment is a structure to store both all a list of points and the Segment that joins
// first and last point
type LineSegment struct {
	Points         []Point4D
	Segment        Segment
	NeedDecimation bool
}

// Line is a structure to store all the LineSegments that compose a line
type Line struct {
	LineSegments []LineSegment
}

func (l Line) NeedDecimation() bool {
	answer := false
	for _, lineSegment := range l.LineSegments{
		if lineSegment.NeedDecimation{
			answer = lineSegment.NeedDecimation
			break
		}
	}
	return answer
}

func (l Line) decimate(threshold float32) Line {
	if !l.NeedDecimation(){
		return l
	}
	return l
}

// ReadFile opens and reads a file and returns its contents
//
// Parameters:
//   - inputFile: path to the input file
//
// Returns:
//   - the content of the input file as an array of bytes
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
//   - inputFile: path to the input file
//
// Returns:
//   - the content of the input file as an array of Point4D
func readPoints(filename string) []Point4D {

	bytes, err := ReadFile(filename)
	var points []Point4D
	err = json.Unmarshal(bytes, &points)
	if err != nil {
		log.Fatalf("Error while unmarshalling JSON: %v", err)
	}
	return points
}

/*
// DecimatePoints remove points that are do not exceed the provided threshold
//
// Parameters:
//	- points: the list of points 4D to decimate
//	- threshold: maximum threshold
//
// Returns:
//	- list of Point4D
func DecimatePoints(listPoints []ListPoints, threshold float32) []Point4D {

}*/

func main() {
	inputFile := "./fixtures/input-4d.json"
	var threshold float32
	threshold = 1.0

	points := readPoints(inputFile)
	segment := Segment{
		Point1: points[0],
		Point2: points[len(points)-1],
	}
	lineSegment := LineSegment{
		Points:         points,
		Segment:        segment,
		NeedDecimation: true,
	}
	line := Line{
		LineSegments: []LineSegment{lineSegment},
	}

	fmt.Println(line.LineSegments[0].Points)
	fmt.Println(line.LineSegments[0].Segment)

	line = line.decimate(threshold)
	fmt.Println(line)
}
