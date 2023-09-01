package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ToBinary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res[0:64]
}

func ReadFile(putanja string) ([]string, []int) {
	f, err := os.Open(putanja)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	words := make([]string, 0)
	weights := make([]int, 0)
	for scanner.Scan() {
		temp := scanner.Text()
		if strings.Contains(temp, ",") || strings.Contains(temp[len(temp)-1:len(temp)], ".") {
			temp = temp[:len(temp)-1]
		}

		postoji := false
		index := 0
		for _, word := range words {
			if word == temp {
				postoji = true
				break
			}
			index++
		}
		if postoji == false {
			words = append(words, temp)
			weights = append(weights, 1)
		} else {
			weights[index]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return words, weights
}

func main() {
	//fmt.Println(ReadFile("C:\\Users\\Stevan\\Desktop\\Projects\\NAISP\\Vezbe4\\SimHash\\tekst1.txt"))
	//fmt.Println(ReadFile("C:\\Users\\Stevan\\Desktop\\Projects\\NAISP\\Vezbe4\\SimHash\\tekst2.txt"))

	words1, weights1 := ReadFile("C:\\Users\\Stevan\\Desktop\\Projects\\NAISP\\Vezbe4\\SimHash\\tekst1.txt")
	words2, weights2 := ReadFile("C:\\Users\\Stevan\\Desktop\\Projects\\NAISP\\Vezbe4\\SimHash\\tekst2.txt")

	sh := newSimHash(words1, weights1, words2, weights2)
	sh.HemingovaUdaljenost()

	//fmt.Println(ToBinary(GetMD5Hash("Sum")))
	//fmt.Println(ToBinary(GetMD5Hash("Hello")))
}
