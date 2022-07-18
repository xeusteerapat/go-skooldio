package main

import "fmt"

func main() {
	var areaOfRec, areaOfTri Object
	areaOfRec = rectangle{width: 4.0, length: 5.0}
	areaOfTri = triangle{width: 4.0, length: 5.0}

	fmt.Println(areaOfRec.Area())
	fmt.Println(areaOfTri.Area())

	// check if value of areaOfTri is type of triangle
	if value, ok := areaOfTri.(triangle); ok {
		value.length = 6
	}
}

type Object interface {
	Area() float64
}

type rectangle struct {
	width, length float64
}

func (rec rectangle) Area() float64 {
	return rec.width * rec.length
}

type triangle struct {
	width, length float64
}

func (tri triangle) Area() float64 {
	return tri.width * tri.length * 0.5
}
