package main

//sample package
import (
	"fmt"
)

func which(util []string) string {
	fmt.Printf("Package : %v \n", util)
	return "utils"
}

// Varoadoc function
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
