// Package decimator provides functionalities to decimate, remove points, from a trajectory
// defined as a set of 4D points.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)

// Point4D is a structure to store spatial/temporal definition of a four dimensional point
type Point4D struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	T float32 `json:"t"`
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

// DouglasPeucker returns a downsampling of a list of Point4D, applying the Ramer–Douglas–Peucker_algorithm
// source: https://en.wikipedia.org/wiki/Ramer%E2%80%93Douglas%E2%80%93Peucker_algorithm
// This function is recursivelly applied over the list of Point4D. Data is stored in place, so only data
// between indexes 0 and index1 are the output of the algorithm.
func DouglasPeucker(points []Point4D, epsilon float32) []Point4D {
	var distance float32
	var distance_maximum float32
	distance_maximum = 0.0

	size_points := len(points)
	if size_points < 2 {
		log.Fatal("Length of point list must be greater than one")
	}

	if size_points == 2 {
		return points
	}

	var index int
	for i := 1; i < size_points; i++ {
		distance = PerpendicularDistance(points[i], points[0], points[size_points-1])
		if distance > distance_maximum {
			index = i
			distance_maximum = distance
		}
	}

	if distance_maximum > epsilon {
		first_slice := DouglasPeucker(points[:index+1], epsilon)
		second_slice := DouglasPeucker(points[index+1:], epsilon)
		points = append(first_slice, second_slice...)
	} else {
		points = append(points[:1], points[len(points)-1])
	}
	return points

}

func PerpendicularDistance(P, A, B Point4D) float32 {
	numerador := math.Abs(float64((B.Y-A.Y)*P.X - (B.X-A.X)*P.Y + B.X*A.Y - B.Y*A.X))
	denominador := math.Sqrt(float64((B.Y-A.Y)*(B.Y-A.Y) + (B.X-A.X)*(B.X-A.X)))
	return float32(numerador / denominador)
}

func main() {
	inputFile := "./fixtures/input-4d.json"
	var threshold float32
	threshold = .50

	points := readPoints(inputFile)
	points = DouglasPeucker(points, threshold)

	fmt.Println(points)
}
