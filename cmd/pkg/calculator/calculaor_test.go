package calculator

import "testing"

func TestSumTwoNums(t *testing.T) {
	result := SumNums(2, 2)
	if result != 4 {
		t.Errorf("Expected 4 and got %f", result)
	}
}

func TestSubtractNums(t *testing.T) {
	result := SubtractNums(2, 2)
	if result != 0 {
		t.Errorf("Expected 0 and got %f", result)
	}
}

func TestMultipleNums(t *testing.T) {
	result := MultipleNums(2, 3)
	if result != 6 {
		t.Errorf("Expected 6 and got %f", result)
	}
}

func TestDivideNums(t *testing.T) {
	result := DivideNums(2, 2)
	if result != 1 {
		t.Errorf("Expected 1 and got %f", result)
	}
}

func TestModuleNums(t *testing.T) {
	result := ModuleNums(3, 2)
	if result != 1 {
		t.Errorf("Expected 1 and got %d", result)
	}
}
