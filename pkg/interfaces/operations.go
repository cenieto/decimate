package interfaces

import "gonum.org/v1/gonum/mat"

type Geometry interface {
	Dimension() int
	CrossProduct(Vector, Vector) Vector
	CrossProductNorm(Vector, Vector) float64
	DoubleAreaTriangle(Point, Line) float64
	DistancePointLine(Point, Line) float64
}

type Point interface {
	mat.Vector
	String() string
	Dimension() int
}

type Line interface {
	String() string
	VectorDirector() Vector
	Length() float64
	Dimension() int
}

type Vector interface {
	mat.Vector
	String() string
	Dimension() int
	Length() float64
	At(int, int) float64
}
