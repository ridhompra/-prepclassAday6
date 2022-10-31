package main

import "fmt"

func main() {
	var input int
	var output string

	fmt.Print("Input\t: ")
	fmt.Scan(&input)
	if input%400 == 0 {
		output = "Kabisat"
	} else if input%100 == 0 {
		output = "Bukan Kabisat"
	} else if input%4 == 0 {
		output = "kabisat"
	} else {
		output = "Bukan Kabisat"
	}

	fmt.Printf("Output\t: %s\n", output)
}
