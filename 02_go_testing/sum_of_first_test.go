package sum_of_first

import "testing"

func TestSumOfFirst(t *testing.T) {
	given := 3
	expect := 6

	get := sumOfFirst(given)

	if expect != get {
		t.Errorf("given %d expect %d but %d\n", given, expect, get)
	}
}
