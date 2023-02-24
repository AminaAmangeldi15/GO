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

func(v *Vector) IsEmty() bool{ 
    if len(v.element) == 0{ 
        fmt.Println("It is empty!") 
        return true; 
    } 
    fmt.Println("It is not empty!") 
    return false; 
} 

func(v *Vector) Contains(element int) bool{ 
    for _,v := range v.element{ 
        if v == element{ 
            fmt.Println("Yes it is found!") 
            return true 
        } 
    } 
    return false 
} 

func (v *Vector) pop_back() (int) { 
    if len(v.element) == 0 { 
        return 0 
    } 
    lastId := len(v.element) - 1 
    lastVal := v.element[lastId] 
    v.element = v.element[:lastId] 
    return lastVal 
} 
 
func (v *Vector) pop_front() (int) { 
    if len(v.element) == 0 { 
        return 0 
    } 
    firstVal := v.element[0] 
    v.element = v.element[1:] 
    return firstVal 
} 