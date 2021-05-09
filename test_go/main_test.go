package test_go 

import (
	// "fmt"
	"testing"
)

type AddResult struct {
	x int
	y int
	expected int

}
var addResults = []AddResult {
	{1, 1, 2},

}
func TestAdd(t *testing.T) {
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if (result != test.expected) {
			t.Fatal("Expected Result Not given")
		}
	}
}
func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("expected 2+2 to equal 4")
	}
}


func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{9999, 10001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

