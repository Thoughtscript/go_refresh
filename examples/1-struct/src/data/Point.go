package data

// Note this is a weakly OOP language
// Note the lack of inheritance and constructors

type Point struct {
	X	float32
	Y	float32
	Z	float32
}

// Pseudo constructor/factory
// Note capitalized name indicates "to export"
func NewPoint(X float32, Y float32, Z float32) *Point{

	// Partially initialize object
	p := Point{X:X, Y: Y}
	p.Z = Z

	// Return pointer
	return &p
}

// Pseudo-class methods
func AddCoordinates(p Point) float32 {
	return p.X + p.Y + p.Z
}