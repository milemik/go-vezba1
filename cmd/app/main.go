package main

import (
	"fmt"
	"log"
	"os"

	"github.com/milemik/go-vezba1/cmd/pkg/calculator"
	"github.com/milemik/go-vezba1/cmd/pkg/checker"
)

func main() {

	inputs := os.Args

	if len(inputs) != 3 {
		log.Println("Please specify only two numbers!")
		os.Exit(0)
	}

	firstNum := inputs[1]
	secondNum := inputs[2]

	firstNumF, secondNumF, err := checker.CheckInputs(firstNum, secondNum)
	if err != nil {
		os.Exit(0)
	}

	fmt.Println(firstNumF, " + ", secondNumF, " is ", calculator.SumNums(firstNumF, secondNumF))
	fmt.Println(firstNumF, " - ", secondNumF, " is ", calculator.SubtractNums(firstNumF, secondNumF))
	fmt.Println(firstNumF, " * ", secondNumF, " is ", calculator.MultipleNums(firstNumF, secondNumF))
	fmt.Println(firstNumF, " / ", secondNumF, " is ", calculator.DivideNums(firstNumF, secondNumF))
	fmt.Println(firstNumF, " % ", secondNumF, " is ", calculator.ModuleNums(firstNumF, secondNumF))
}
