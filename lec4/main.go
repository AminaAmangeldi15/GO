package main

import (
	"fmt"
	ar "github.com/AminaAmangeldi15/Go/lec4/hw"
)

func main() {
	a := ar.ArrayList{}
	a.Add(5)
	a.Add(6)
	a.Size()
	a.Remove(0)
	a.Size()
	fmt.Println(a.Get(1))
	fmt.Println(a)
	v := ar.Vector{}
	v.Add(3)
	v.Size()
}