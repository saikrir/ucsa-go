package main

import "fmt"

func main() {
	var from, to, by int32 = 0, 250, 10
	var farenheit = from
	for farenheit <= to {
		centigrade := 5.0 / 9.0 * float32(farenheit-32)
		fmt.Printf("Farentheit %.2f is %.2f Cenigrade \n", float32(farenheit), centigrade)
		farenheit += by
	}
}
