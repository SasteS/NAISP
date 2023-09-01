package main

import "fmt"

type CountMinSketch struct {
	m               uint
	k               uint
	set             [][]byte
	nizHashFunkcija []HashWithSeed
}

func (cms CountMinSketch) DoAllHashes(s string) []uint64 {
	//k_ := bf.k
	m_ := cms.m
	//copy_set := bf.set
	copyhash := cms.nizHashFunkcija

	data := []byte(s)

	niz_hesiranih_vrednosti := make([]uint64, 0)
	for _, fn := range copyhash {
		niz_hesiranih_vrednosti = append(niz_hesiranih_vrednosti, fn.Hash(data)%uint64(m_))
	}
	fmt.Println(niz_hesiranih_vrednosti)
	return niz_hesiranih_vrednosti
}

func (cms CountMinSketch) UpdateGrid(s string) {
	set_copy := cms.set

	niz_hesiranih_vrednosti := cms.DoAllHashes(s)

	for i := 0; i < len(niz_hesiranih_vrednosti); i++ {
		set_copy[i][niz_hesiranih_vrednosti[i]]++
	}

	cms.set = set_copy
	fmt.Println(cms.set)
}

func (cms CountMinSketch) Check(s string) {
	set_copy := cms.set

	niz_hesiranih_vrednosti := cms.DoAllHashes(s)

	checked_values := make([]byte, len(niz_hesiranih_vrednosti))
	for i := 0; i < len(niz_hesiranih_vrednosti); i++ {
		checked_values[i] = set_copy[i][niz_hesiranih_vrednosti[i]]
	}

	var min byte = 255
	for _, item := range checked_values {
		if min > item {
			min = item
		}
	}
	fmt.Print("Ucestalost: ")
	fmt.Print(min)
}

func newCountMinSketch() *CountMinSketch {
	var m_ uint = CalculateM(0.001)
	var k_ uint = CalculateK(0.999)

	set_ := make([][]byte, k_)
	for i := 0; i < int(k_); i++ {
		set_[i] = make([]byte, m_)
	}
	//fmt.Println(set_)

	nizHashFunkcija_ := CreateHashFunctions(k_)

	cms := CountMinSketch{set: set_, nizHashFunkcija: nizHashFunkcija_, k: k_, m: m_}
	return &cms
}
