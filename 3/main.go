package main

import "fmt"

func main() {
	var input int
	var output string

	fmt.Print("Input\t: ")
	fmt.Scan(&input)

	if input%2 == 0 {
		output = "Genap"
	} else {
		output = "ganjil"
	}
	fmt.Printf("Output\t: %s\n", output)

}
