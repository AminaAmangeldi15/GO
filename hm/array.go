package hm

import (
	"fmt"
	"strconv"
)

type ArrayList struct {
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
