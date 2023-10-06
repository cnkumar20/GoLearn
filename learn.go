package main

import (
	"fmt"
)

func daysOfWork() (x, y string, err error) {
	x = "Monday"
	y = "Thursday"
	err = nil
	return
}
func getlength(name, color string) int {
	var n = len(name)
	var c = len(color)
	return c + n
}
func modifyMap(names []string, i int) []string {
	fmt.Printf("Before: %v \n", names)
	names = append(names[:i], names[i+1:]...)
	fmt.Printf("After: %v \n", names)
	return names
}

type student struct {
	name string
	age  int
}

func main() {
	var num int = 5
	const name = "kumar"
	fmt.Println("len:", num, " name:", name)
	var fmt_name = fmt.Sprintf(name)
	fmt.Println("name:", fmt_name)
	msg := 0.55
	fmt.Printf("%.2f \n", msg)
	if 3 < 3 {
		fmt.Print("2 < 3\n")
	} else if 3 < 4 {
		fmt.Print("3<4\n")
	} else {
		fmt.Printf("None \n")
	}
	if length := getlength("Nandiesh", "brown"); length > 1 {
		fmt.Printf("this is length %d \n", length)
	}
	x, y, err := daysOfWork()
	if err != nil {
		fmt.Printf("%s , %s \n", x, y)
	}
	var s1 student
	s1.name = "nandiesh"
	s1.age = 35
	fmt.Printf("Name : %s \n Age: %d\n", s1.name, s1.age)

	type Rectangle struct {
		width  int
		height int
	}
	check1 := func(name string) int {
		i := len(name)
		return i
	}
	fmt.Printf("Length of name :%d \n", check1("Kumar"))
	slice_ex := make([]int, 6)
	fmt.Println(slice_ex[:6])
	slice_ex = append(slice_ex, 6)
	fmt.Println(len(slice_ex))
	fmt.Println(slice_ex[2:])
	n2 := map[string]int{"foo": 1, "bar": 2}
	n2["kumar"] = 3
	fmt.Println(n2["kumar"])
	counts := make(map[string]int)
	counts["one"] = 1
	fmt.Printf("One: %d\n", counts["one"])
	for x := 0; x < 2; x++ {
		c := ""
		if x == 0 {
			c = "one"
		} else if x == 1 {
			c = "two"
		}
		elem, err := counts[c]
		if err != true {
			fmt.Println("Element doesn't exist")
		}
		fmt.Printf("Iter elem : %d \n", elem)

	}

	//Pointers
	day := "Thursday"
	var holiday *string
	holiday = &day
	fmt.Printf("Today is : %s \n", day)

	fmt.Printf("Today holiday print mem address : %s \n", holiday)
	fmt.Printf("Today holiday : %s \n", *holiday)
	//modify underneat
	*holiday = "Friday"
	fmt.Printf("Today is : %s \n", day)

	fmt.Printf("Today holiday print mem address : %s \n", holiday)
	fmt.Printf("Today holiday : %s \n", *holiday)

	//slice
	names := make([]string, 0)
	names = append(names, "kumar", "nand", "priya", "avani", "google")
	fmt.Printf("Outside Before: %v \n", names)
	modifyMap(names, 3)
	fmt.Printf("Outside After: %v \n", names)

}
