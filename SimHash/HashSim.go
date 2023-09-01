package main

import (
	"fmt"
	"strconv"
)

type SimHash struct {
	uredjeni1           []int
	uredjeni2           []int
	hemingovaUdaljenost int
}

func (sh SimHash) HemingovaUdaljenost() {
	//bin_uredjeni1 := make([]byte, 0)
	//bin_uredjeni2 := make([]byte, 0)

	for i := 0; i < len(sh.uredjeni1); i++ {
		if sh.uredjeni1[i] > 0 {
			sh.uredjeni1[i] = 1
		} else {
			sh.uredjeni1[i] = 0
		}
	}

	for i := 0; i < len(sh.uredjeni2); i++ {
		if sh.uredjeni2[i] > 0 {
			sh.uredjeni2[i] = 1
		} else {
			sh.uredjeni2[i] = 0
		}
	}

	s := ""
	for i := 0; i < len(sh.uredjeni2); i++ {
		if sh.uredjeni1[i]+sh.uredjeni2[i] == 1 {
			s += "1"
		} else {
			s += "0"
		}
	}

	fmt.Println(sh.uredjeni1, sh.uredjeni2)
	fmt.Println(len(sh.uredjeni1), len(sh.uredjeni2))
	fmt.Println(s)
	fmt.Println(strconv.ParseUint(s, 2, 64))
}

func newSimHash(words1 []string, weights1 []int, words2 []string, weights2 []int) *SimHash {
	uredjeni1_ := make([]int, 0)
	uredjeni2_ := make([]int, 0)

	hashWords1 := make([]string, 0)
	for _, word := range words1 {
		hashWords1 = append(hashWords1, ToBinary(GetMD5Hash(word)))
	}

	hashWords2 := make([]string, 0)
	for _, word := range words2 {
		hashWords2 = append(hashWords2, ToBinary(GetMD5Hash(word)))
	}

	for i := 0; i < 64; i++ {
		temp_sum := 0
		for j := 0; j < len(hashWords1); j++ {
			word := hashWords1[j]
			if word[i:i+1] == "0" {
				temp_sum += weights1[j] * -1
			} else {
				temp_sum += weights1[j] * 1
			}
		}
		uredjeni1_ = append(uredjeni1_, temp_sum)

		temp_sum = 0
		for j := 0; j < len(hashWords2); j++ {
			word := hashWords2[j]
			if word[i:i+1] == "0" {
				temp_sum += weights2[j] * -1
			} else {
				temp_sum += weights2[j] * 1
			}
		}
		uredjeni2_ = append(uredjeni2_, temp_sum)
	}

	sh := SimHash{uredjeni1: uredjeni1_, uredjeni2: uredjeni2_, hemingovaUdaljenost: 0}

	return &sh
}
