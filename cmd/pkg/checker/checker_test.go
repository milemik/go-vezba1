package checker

import (
	"testing"
)

func TestCheckerInt(t *testing.T) {
	first, second, _ := CheckInputs("12", "12")

	if first != 12 {
		t.Errorf("Expected 12 and got %v", first)
	}

	if second != 12 {
		t.Errorf("Expected 12 and got %v", second)
	}
}

func TestCheckerFloat(t *testing.T) {
	first, second, _ := CheckInputs("12.12", "12.15")

	if first != 12.12 {
		t.Errorf("Expected 12.12 and got %v", first)
	}

	if second != 12.15 {
		t.Errorf("Expected 12.15 and got %v", second)
	}
}

func TestCheckerFirstString(t *testing.T) {
	_, _, err := CheckInputs("abc", "12.15")
	if err == nil {
		t.Errorf("Expected error here... first value is string")
	}
}

func TestCheckerSecondString(t *testing.T) {
	_, _, err := CheckInputs("12.15", "vd")
	if err == nil {
		t.Errorf("Expected error here... first value is string")
	}
}

// func TestConfertToFloat(t *testing.T) {
// 	val, err := convertToFloat("12")

// 	if err != nil {
// 		t.Error("Didt expect error here")
// 	}
// 	if val != 12 {
// 		t.Errorf("Expected 12 and got %f", val)
// 	}
// }
