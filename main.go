package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Point4D struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	T float32 `json:"t"`
}

type Segment struct {
	Point1 Point4D
	Point2 Point4D
}

func readFile(inputFile string) ([]byte, error) {

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
func readPoints(filename string) []Point4D {

	bytes, err := readFile(filename)
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
