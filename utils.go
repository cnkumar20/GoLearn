package main

//sample package
import (
	"fmt"
)

func which(util []string) string {
	fmt.Printf("Package : %v \n", util)
	return "utils"
}

// Variadic function
func workingDay(days ...string) (numDays int, err error) {
	numDays = len(days)
	fmt.Printf("Total number of days %d \n", len(days))
	for i, day := range days {
		fmt.Printf("What day is it on %d : %s \n", i, day)
	}
	if numDays > 7 {
		err = fmt.Errorf("Number of days can't be more than 7")
	}
	return numDays, err
}

// Closure Function / Anonymous Function
func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fib(num int) int {
	if num == 0 {
		return 1
	}
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 1
	}
	return fib(num-1) + fib(num-2)

}

func modifyPointerVariable(i *int) *int {
	//increment
	*i = *i + 1
	return i
}
