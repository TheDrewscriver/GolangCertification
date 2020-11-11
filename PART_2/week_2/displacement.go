package main

import (
	"fmt"
	"math"
	"os"
)

/**
This function validates the input. If there is an error,
we exit the program
*/
func validateInput(errorValue error) {
	if errorValue != nil {
		fmt.Println("Can only accept floatValues for input")
		os.Exit(1)
	}
}

func main() {
	/*
		The variables we use
	*/
	var a float64
	var v0 float64
	var s0 float64
	var t float64

	//Get all the input values
	fmt.Println("Enter Value for Acceleration(a)")
	_, error1 := fmt.Scanf("%f", &a)
	validateInput(error1)

	fmt.Println("Enter Value for Initial Velocity(v0)")
	_, error2 := fmt.Scanf("%f", &v0)
	validateInput(error2)

	fmt.Println("Enter Value for Initial Displacement(s0)")
	_, error3 := fmt.Scanf("%f", &s0)
	validateInput(error3)

	fmt.Println("Enter Value for Time(t)")
	_, error4 := fmt.Scanf("%f", &t)
	validateInput(error4)

	//Get the returned function
	funcVar := GenDisplaceFn(a, v0, s0)
	//Print the displacement
	fmt.Printf("Displacement : %f \n", funcVar(t))
}

//GenDisplaceFn Function that returns another function that we use to generate the
//displacement
func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	var function = func(t float64) float64 {
		var t0 = float64(t)
		return (0.5 * a * math.Pow(t0, 2)) + (v0 * t0) + s0
	}
	return function
}
