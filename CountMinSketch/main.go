package main

import "fmt"

func main() {
	cms := newCountMinSketch()
	fmt.Println(cms.m)
	fmt.Println(cms.k)
	cms.UpdateGrid("Asdw")
	cms.Check("Asdw")
}
