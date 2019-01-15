package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var alphabet = strings.Split("A-B-C-D-E-F-G-H-I-J-K-L-M-N-O-P-Q-R-S-T-U-V-W-X-Y-Z", "-")

func Task22(filename string) int {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	totalScore := 0
	data := strings.Replace(string(dat), "\"", "", -1)
	names := strings.Split(data, ",")
	sort.Strings(names)
	for idx, name := range names {
		score := CalculateScore(idx, name)
		totalScore += score
	}
	return totalScore
}

func CalculateScore(idx int, name string) int {
	score := 0
	chars := strings.Split(name, "")
	for _, char := range chars {
		score += GetAlphabetPosition(char)
	}
	score *= (idx + 1)
	return score
}

func GetAlphabetPosition(char string) int {
	for idx, c := range alphabet {
		if char == c {
			return idx + 1
		}
	}
	panic(char)
}
