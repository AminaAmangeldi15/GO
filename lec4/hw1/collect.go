package main

import (
	"fmt"
	"strconv"
	// "strings"
)

type List interface {
	Add(element int)
	Remove(index int) int
	Get(index int) int
	Size() int
}

type ArrayList struct {
	element []int
}

type Vector struct {
	element []int
}

func (a *ArrayList) Add(element int) {
	a.element = append(a.element, element)
	fmt.Println("Element added to the array")
}

func (a *ArrayList) Remove(index int) int {
	if index < 0 || index >= len(a.element) {
		fmt.Println("Index out of range")
		return 0
	}
	removedElement := a.element[index]
	a.element = append(a.element[:index], a.element[index+1:]...)
	fmt.Println("Element removed from the array")
	return removedElement
}

func (a *ArrayList) Get(index int) int {
	if index < 0 || index >= len(a.element) {
		fmt.Println("Index out of range")
		return 0
	}
	return a.element[index]
}

func (a *ArrayList) Size() {
	size := strconv.Itoa(len(a.element))
	fmt.Println("Size of array " + size)
}


func (v *Vector) Add(element int) {
	v.element = append(v.element, element)
	fmt.Println("Element added to the vector")
}

func (v *Vector) Remove(index int) int {
	if index < 0 || index >= len(v.element) {
		fmt.Print("Index out of range ")
		return 0
	}
	removedElement := v.element[index]
	v.element = append(v.element[:index], v.element[index+1:]...)
	fmt.Println("Element removed from the vector")
	return removedElement
}

func (v *Vector) Get(index int) int {
	if index < 0 || index >= len(v.element) {
		fmt.Print("Index out of range ")
		return 0
	}
	return v.element[index]
}

func (v *Vector) Size() {
	size := strconv.Itoa(len(v.element))
	fmt.Println("Size of vector " + size)
}

func main() {
	a := ArrayList{}
	a.Add(5)
	a.Add(6)
	a.Size()
	a.Remove(0)
	a.Size()
	fmt.Println(a.Get(1))
	// fmt.Println(a)
	// v := Vector{}
	// v.Add(3)
	// v.Size()
}
