package AbstractFactory
// not good to use with Golang
type (
	Shape interface {
		Draw() float64
	}
	Color	interface {
		Fill(ColorName)
	}
	Rectangle struct {
		Width 	float64
		Height 	float64
	}
	Square	struct {
		Edge	float64
	}
	Circle	struct {
		Radius 		float64
	}

	ColorName	string
)
