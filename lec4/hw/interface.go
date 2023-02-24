package hw

type List interface {
	Add(element int)
	Remove(index int) int
	Get(index int) int
	Size() int
	isEmpty() bool 
    Contains(element int) bool 
}