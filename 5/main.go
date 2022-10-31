package main

import "fmt"

func main() {
	var Input int
	var output string

	fmt.Print("Input\t: ")
	fmt.Scan(&Input)

	usia := condition(Input, output)

	fmt.Printf("Output\t: %s\n", usia)
}
func condition(input int, output string) string {
	output = "Semua usia"
	if input >= 21 {
		output = "Dewasa"
	}
	if input >= 13 && input < 21 {
		output = "Remaja"
	}
	if input >= 9 && input < 13 {
		output = "Bimbingan orang tua"
	}
	return output
}
