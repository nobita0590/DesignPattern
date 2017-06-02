package FactoryPattern
// not good to use with Golang
import "math"

type (
	Shape interface {
		Draw() float64
	}

	Rectangle	struct {
		Width 		float64
		Height		float64
	}

	Circle		struct {
		Radius 		float64
	}

	ShapeFactory	struct {

	}
	ShapeName	string
)

const (
	CircleName ShapeName = "Circle"
	RectangleName ShapeName = "Rectangle"
)

func (r Rectangle) Draw() float64 {
	return 2*(r.Width + r.Height)
}

func (c Circle) Draw() float64 {
	return c.Radius * 2 * math.Pi
}

func (sf ShapeFactory) GetShape(shapeType ShapeName) Shape {
	switch shapeType {
	case CircleName:
		return Circle{}
	case RectangleName:
		return Rectangle{}
	}
	return nil
}
