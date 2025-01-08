package geom2d

type Segment2D struct {
	PointA Point2D
	PointB Point2D
}

func NewSegment(pa, pb Point2D) Segment2D {
	return Segment2D{
		PointA: pa,
		PointB: pb,
	}
}

func (s Segment2D) Dimension() int {
	return 2
}
