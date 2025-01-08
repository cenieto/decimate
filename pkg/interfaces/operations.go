package interfaces

type Vectorial interface {
	Dimension() int
	Components() []float64
	CrossProduct(other Vectorial) (Vectorial, error)
}

type Segment interface {
	Dimension() int
	TriangleDefinedArea() int
}
