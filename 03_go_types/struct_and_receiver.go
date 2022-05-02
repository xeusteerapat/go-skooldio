package main

import "fmt"

type Rectangle struct {
	w, l float64
}

// function signature that
func Area(rec Rectangle) float64 {
	return rec.w * rec.l
}

// receiver function AKA object method
func (rec Rectangle) getArea() float64 {
	return rec.w * rec.l
}

func main() {
	myRec := Rectangle{
		w: 4.5,
		l: 6.6,
	}

	fmt.Println(Area(myRec)) // 29.7
	fmt.Println("<-----//----->")
	fmt.Println(myRec.getArea()) // 29.7
}
