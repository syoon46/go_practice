package main

import "fmt"

type shape interface {
    area() int
}
 
type square struct {
    sideLength int
}
 
func (s square) area() int {
    return 10
}
 
func printArea(s shape) {
    fmt.Println(s.area())
}

func main() {
	s1 := square {}

	printArea(s1)
}