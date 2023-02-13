package main

import (
	"fmt"
	"github.com/AminaAmangeldi15/Go/hm"
)

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