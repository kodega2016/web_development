package main

import (
	"testing"
)

func TestDivideBy(
	t *testing.T,
) {
	res, _ := divideBy(20, 10)
	if res != 2 {
		t.Errorf("expected result:2 but got %v", res)
	}

	_, err := divideBy(20, 0)
	if err == nil {
		t.Errorf("expected an error but got the result")
	}
}
