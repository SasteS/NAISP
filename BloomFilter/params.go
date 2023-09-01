package main

import "math"

func CalculateM(expectedElements int, falsePositiveRate float64) uint64 { //msm da je expectedElements velicina koju ocemo a ovaj falsePositiveRate nz
	return uint64(math.Ceil(float64(expectedElements) * math.Abs(math.Log(falsePositiveRate)) / math.Pow(math.Log(2), float64(2))))
}

func CalculateK(expectedElements int, m uint64) uint64 {
	return uint64(math.Ceil((float64(m) / float64(expectedElements)) * math.Log(2)))
}
