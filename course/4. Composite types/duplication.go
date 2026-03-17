package main

import "fmt"

func RemoveDuplicates(input []string) []string {
	var output []string
	var double = map[string]bool{}

	for _, a := range input {
		if _, ok := double[a]; !ok {
			double[a] = true
			output = append(output, a)
		}
	}
	return output
}

func test_duplication() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(RemoveDuplicates(input))
}
