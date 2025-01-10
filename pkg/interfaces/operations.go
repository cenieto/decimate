package interfaces

type Vectorial interface {
	Dimension() int
	Components() []float64
}

type Segment interface {
	Dimension() int
	Components() []Vectorial
}

type Geometry interface {
	Dimension() int
	CrossProduct(Vectorial, Vectorial) Vectorial
}
