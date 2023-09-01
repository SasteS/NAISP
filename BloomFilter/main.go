package main

import "fmt"

func main() {
	bf := newBloom()
	bf.UpdateFilter("String")
	bf.UpdateFilter("Hello")
	bf.UpdateFilter("Yahoo")
	bf.UpdateFilter("Yaho")
	bf.UpdateFilter("Ellp")
	bf.UpdateFilter("a")
	bf.UpdateFilter("i")

	fmt.Println(bf.set)
	bf.Check("Hallo")
	fmt.Println(bf.DoAllHashes("Hallo"))

	/*fns := CreateHashFunctions(5)

	buf := &bytes.Buffer{}
	encoder := gob.NewEncoder(buf)
	decoder := gob.NewDecoder(buf)

	for _, fn := range fns {
		data := []byte("hello") //niz ascii vrednosti
		fmt.Println(data)
		fmt.Println(fn.Hash(data))
		err := encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		fmt.Println(dfn.Hash(data))
	}*/

}
