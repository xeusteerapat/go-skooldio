package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
the default Error interface in Go
type error interface {
	Error() string
}
*/

func handler(s fmt.Stringer) error {
	if s == nil {
		return errors.New("s is nil")
	}

	n, err := strconv.Atoi(s.String())

	if err != nil {
		return err
	}

	_ = n

	return nil
}

func main() {

}
