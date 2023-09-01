package main

import (
	"crypto/md5"
	"encoding/binary"
	"math"
)

const (
	HLL_MIN_PRECISION = 4
	HLL_MAX_PRECISION = 16
)

type HLL struct {
	m   uint64
	p   uint8
	reg []uint8
}

func (hll *HLL) Estimate() float64 {
	sum := 0.0
	for _, val := range hll.reg {
		sum += math.Pow(math.Pow(2.0, float64(val)), -1)
	}

	alpha := 0.7213 / (1.0 + 1.079/float64(hll.m))
	estimation := alpha * math.Pow(float64(hll.m), 2.0) / sum
	emptyRegs := hll.emptyCount()
	if estimation <= 2.5*float64(hll.m) { // do small range correction
		if emptyRegs > 0 {
			estimation = float64(hll.m) * math.Log(float64(hll.m)/float64(emptyRegs))
		}
	} else if estimation > 1/30.0*math.Pow(2.0, 32.0) { // do large range correction
		estimation = -math.Pow(2.0, 32.0) * math.Log(1.0-estimation/math.Pow(2.0, 32.0))
	}
	return estimation
}

func (hll *HLL) emptyCount() int {
	sum := 0
	for _, val := range hll.reg {
		if val == 0 {
			sum++
		}
	}
	return sum
}

func (hll *HLL) addElem(s string) {
	fn := md5.New()
	fn.Write([]byte(s))
	hash := binary.BigEndian.Uint64(fn.Sum(nil))
	i := hash >> (64 - hll.p)
	//fmt.Println(i)
	brNula := 0
	j := 1
	for k := 63; k >= 0; k-- {
		j = int(hash >> uint64(k))
		if j == 1 {
			break
		}
		brNula++
	}
	hll.reg[i] = uint8(brNula)
}

func newHLL(p_ uint8) *HLL {
	m_ := math.Pow(2, float64(p_))

	reg_ := make([]uint8, int(m_))

	hll := HLL{m: uint64(m_), p: p_, reg: reg_}

	return &hll
}
