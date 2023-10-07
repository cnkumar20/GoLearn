package main

//sample package
import (
	"fmt"
)

func which(util []string) string {
	fmt.Printf("Package : %v \n", util)
	return "utils"
}
