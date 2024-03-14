package main

import (
	"fmt"
	"log"
	"math"
)

const pi = 3.14159

func main() {
	var radius, area float64

	fmt.Print("Please Enter Radius of Circle: ")

	_, err := fmt.Scanf("%f", &radius)

	if err != nil {
		log.Fatalln("Invalid input for radius")
	}

	area = pi * math.Pow(radius, 2)

	fmt.Printf("Area of a circle is %4.2f \n", area)
}
