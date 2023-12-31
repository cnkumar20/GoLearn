package main

//sample package
import (
	"fmt"
	"time"
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

func modifySlice(names []string) []string {
	tempArray := &names
	(*tempArray)[1] = "Check"
	return names
}
func modifyArray(names [3]string) [3]string {
	tempArray := &names
	(*tempArray)[1] = "Check"
	fmt.Printf("Inside Modified array : %v \n", *tempArray)
	return names
}

func funcArrayModify(names *[3]string) *[3]string {
	names[1] = "Check"
	fmt.Printf("Inside Modified array : %v \n", *names)
	return names
}

//string learning functions

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

//Generics and Maps

func makeMap[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func channelSyncExample(done chan bool, msg string) {
	fmt.Printf("Inside ChannelSyncExample : %s \n", msg)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Done \n")
	done <- true
}

func ping(msg string, pings chan<- string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
