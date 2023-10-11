package tutor

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func LearnArrays(arr []int) {
	for x := 0; x < len(arr); x++ {
		fmt.Println(arr[x])
	}
}

func PrintArrayUsingRange(arr []int) {
	for i, num := range arr {
		if num == 2 {
			fmt.Printf("if num==2 value == %d and at index = %d \n", num, i)
		}
	}
}

func GoReturnMultipleValues(arr []string) (l int, err error) {
	l = len(arr)
	if l == 0 {
		return 0, fmt.Errorf("No elements in array")
	}
	return l, err
}
