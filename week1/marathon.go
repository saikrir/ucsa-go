package main

import "fmt"

func main() {
	var miles, yards int32 = 26, 385
	var distanceInKm float32 = (1.60934 * (float32(miles) + float32(yards)/1760.0))
	fmt.Printf("Distance in KM %f \n ", distanceInKm)
}
