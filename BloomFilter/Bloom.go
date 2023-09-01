package main

import "fmt"

type Bloom struct {
	k               uint64
	m               uint64
	set             []byte
	nizHashFunkcija []HashWithSeed
}

func (bf Bloom) DoAllHashes(s string) []uint64 {
	//k_ := bf.k
	m_ := bf.m
	//copy_set := bf.set
	copyhash := bf.nizHashFunkcija

	data := []byte(s)

	niz_hesiranih_vrednosti := make([]uint64, 0)
	for _, fn := range copyhash {
		niz_hesiranih_vrednosti = append(niz_hesiranih_vrednosti, fn.Hash(data)%m_)
	}
	//fmt.Println(niz_hesiranih_vrednosti)
	return niz_hesiranih_vrednosti
}

func (bf Bloom) UpdateFilter(s string) {
	set_copy := bf.set

	niz_hesiranih_vrednosti := bf.DoAllHashes(s)
	//fmt.Println(niz_hesiranih_vrednosti)

	for _, item := range niz_hesiranih_vrednosti {
		if set_copy[item] == 0 {
			set_copy[item] = 1
		}
	}

	bf.set = set_copy
	//fmt.Println(bf.set)
}

func (bf Bloom) Check(s string) {
	set_copy := bf.set
	niz_hesiranih_vrednosti := bf.DoAllHashes(s)

	postoji := true
	for _, item := range niz_hesiranih_vrednosti {
		if set_copy[item] == 0 {
			postoji = false
			break
		}
	}

	if postoji == false {
		fmt.Println("no")
	} else {
		fmt.Println("maybe")
	}
}

func newBloom() *Bloom {
	var m_ uint64 = CalculateM(200, 0.0001)
	var k_ uint64 = CalculateK(200, m_)
	tempSet := make([]byte, m_) //nekako moram da namestim oavj falsepositive p
	tempNizHashFja := CreateHashFunctions(k_)

	/*fmt.Println("m = ")
	fmt.Println(m) //velicina bafera
	fmt.Println("k = ")
	fmt.Println(k) //velicina niza hash fja*/

	p := Bloom{set: tempSet, nizHashFunkcija: tempNizHashFja, k: k_, m: m_}

	return &p
}

/*func indexOf(word uint64, data []uint64) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}*/
