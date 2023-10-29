package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func (s *student) print() {
	fmt.Printf("Name : %s \n", s.name)
	fmt.Printf("Age : %d \n", s.age)
}

func (s *student) checkAge() {
	if s.age > 18 {
		fmt.Printf("Student %s is an Adult", s.name)
	}
}

func (s *student) addAge(add int) (age int, err error) {
	if s.age > 0 {
		s.age = s.age + add
	} else {
		fmt.Errorf("This student has invalid age")
	}
	return s.age, err
}

type College interface {
	level()
	course()
}

func (s student) level() {
	if s.age > 16 && s.age < 20 {
		fmt.Printf("High School \n")
	} else if s.age <= 16 {
		fmt.Printf("Not HighSchool \n")
	} else {
		fmt.Printf("Invalid Student \n")
	}
}
func (s student) course() {
	if s.age > 16 {
		fmt.Printf("studied science \n")
	}
}

func printstructinfo(s College) {
	s.level()
	s.course()
}

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b *base) describe() {
	fmt.Printf("Describing the base : %d \n", b.num)
}

type describer interface {
	describe()
}
