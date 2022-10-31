package main

import "fmt"

func main() {
	var input int
	var output string

	fmt.Print("Input\t: ")
	fmt.Scan(&input)

	output = "ERROR INPUT"
	if input >= 90 && input <= 100 {
		output = "A"
	} else if input >= 80 && input < 90 {
		output = "B"
	} else if input >= 70 && input < 80 {
		output = "C"
	} else if input >= 60 && input < 70 {
		output = "D"
	} else if input < 60 {
		output = "E"
	}

	fmt.Printf("Output\t: %s\n", output)

}
