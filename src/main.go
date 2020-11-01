package main

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Line []Line
}

// supplied interface
func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		Line{0, 0, width, 0},
		Line{0, 0, 0, height},
		Line{width, 0, width, height},
		Line{0, height, width, height},
	}}
}

// interface you currently have
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	//
	return ""
}

func main() {

}
