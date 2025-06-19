package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		str string
	}{
		{"([])"},
		{"(]"},
		{"){"},
		{"(){}}{"},
		{"()"},
	}
	for _, c := range cases {
		fmt.Println(c.str)
		fmt.Println(isValid(c.str))
	}
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		str string
	}{
		{"3+2*2"},
	}
	for _, c := range cases {
		fmt.Println(c.str)
		fmt.Println(calculate(c.str))
	}
}
