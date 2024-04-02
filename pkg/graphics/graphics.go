package graphics

type Point2D struct {
	x int
	y int
}

func NewPoint(x int, y int) Point2D {
	return Point2D{x: x, y: y}
}

func (point Point2D) Coordinates() (int, int) {
	return point.x, point.y
}

// source: https://github.com/StephaneBunel/bresenham/blob/master/drawline.go
func GetLinePoints(p1 Point2D, p2 Point2D) (linePoints []Point2D) {
	linePoints = make([]Point2D, 0)

	var x1, y1 = p1.Coordinates()
	var x2, y2 = p2.Coordinates()

	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		linePoints = append(linePoints, Point2D{x: x1, y: y1})

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			linePoints = append(linePoints, Point2D{x: x1, y: y1})
			x1++
		}
		linePoints = append(linePoints, Point2D{x: x1, y: y1})

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1 = y2
		}
		for ; dy != 0; dy-- {
			linePoints = append(linePoints, Point2D{x: x1, y: y1})
			y1++
		}
		linePoints = append(linePoints, Point2D{x: x1, y: y1})

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				linePoints = append(linePoints, Point2D{x: x1, y: y1})
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				linePoints = append(linePoints, Point2D{x: x1, y: y1})
				x1++
				y1--
			}
		}
		linePoints = append(linePoints, Point2D{x: x1, y: y1})

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				linePoints = append(linePoints, Point2D{x: x1, y: y1})
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				linePoints = append(linePoints, Point2D{x: x1, y: y1})
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		linePoints = append(linePoints, Point2D{x: x2, y: y2})

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				linePoints = append(linePoints, Point2D{x: x1, y: y1})
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				linePoints = append(linePoints, Point2D{x: x1, y: y1})
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		linePoints = append(linePoints, Point2D{x: x2, y: y2})
	}
	return
}
