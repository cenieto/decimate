package interfaces

type Geometry interface {
	Dimension() int
	CrossProduct(Vectorial, Vectorial) Vectorial
}
