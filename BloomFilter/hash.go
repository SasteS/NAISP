package main

import (
	"crypto/md5"
	"encoding/binary"
	"time"
)

type HashWithSeed struct {
	Seed []byte
}

func (h HashWithSeed) Hash(data []byte) uint64 { //deluje kao da je ovo fja od strukture
	fn := md5.New()
	fn.Write(append(data, h.Seed...))
	return binary.BigEndian.Uint64(fn.Sum(nil))
}

func CreateHashFunctions(k uint64) []HashWithSeed {
	h := make([]HashWithSeed, k)  //napravi 5 hashFunkcija? //mislim da ce zauzeti memorije koliko mu treba za niz tih 5 hashova i prima samo HashWithSeed objekte
	ts := uint(time.Now().Unix()) //vrednost trenutnog vremena izrazeno u sek vljd
	for i := uint(0); i < uint(k); i++ {
		seed := make([]byte, 32)                       //prazan niz koji prima 32 byte elementa?
		binary.BigEndian.PutUint32(seed, uint32(ts+i)) //izgleda da napuni ceo niz enkodiranim vrednostima?
		hfn := HashWithSeed{Seed: seed}                //pretpostavljam dodela seed tog nekog niza Seed nizu iz strukture? msm da je to fora s konstruktorom da se dodeljuju vrednosti Seedu
		h[i] = hfn                                     //samo ide redom kroz niz hashFja i dodeljuje ovu hfn kao clan niza
	}
	return h
}
