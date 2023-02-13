package main

import (
	"fmt"
	a "github.com/AminaAmangeldi15/Go/lec4/hw"
)

func main() {
	a := a.ArrayList{}
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