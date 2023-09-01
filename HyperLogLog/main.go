package main

import "fmt"

func main() {
	hll := newHLL(4)
	fmt.Println(hll)
	hll.addElem("String")
	hll.addElem("Awdsda")
	hll.addElem("kefeENLSNEFNS")
	fmt.Println(hll.Estimate())
}
