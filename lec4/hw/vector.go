package hw

import (
	"fmt"
	"strconv"
)

type Vector struct {
	element []int
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