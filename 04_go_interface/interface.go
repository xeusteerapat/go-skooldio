package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type String string

func (s String) String() string {
	return string(s)
}

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	ans := Int(4)

	fmt.Println(ans)
}
