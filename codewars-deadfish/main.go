package main

import "fmt"

func deadfish_interpreter(deadfish_input string) []int {

	var current_value int = 0
	var output []int

	for i := 0; i < len(deadfish_input); i++ {
		switch deadfish_input[i] {
		case 'i':
			current_value++
		case 'd':
			current_value--
		case 's':
			current_value = current_value * current_value
		case 'o':
			output = append(output, current_value)

		}
	}
	return output
}

func main() {
	fmt.Println(deadfish_interpreter("iiisdoso"))
	fmt.Println(deadfish_interpreter("iiisdosodddddiso"))
}
