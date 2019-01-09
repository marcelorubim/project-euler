package main

import (
	"errors"
	"math"
	"strings"
)

func Task17(lim int) int {
	count := 0
	for i := 1; i <= lim; i++ {
		count += len(strings.Replace(strings.Replace(ConvertNumberToText(i), " ", "", -1), "-", "", -1))
	}
	return count
}

func ConvertNumberToText(num int) string {
	s := ""
	space := ""
	if num >= 1000 && num < 10000 {
		text, _ := GetNumberTextUpToTen(int(math.Floor(float64(num) / 1000)))
		s = strings.Join([]string{s, text}, "")
		s = strings.Join([]string{s, "thousand"}, " ")
		space = " "
		num = num % 1000
	}
	if num >= 100 && num < 1000 {
		text, _ := GetNumberTextUpToTen(int(math.Floor(float64(num) / 100)))
		s = strings.Join([]string{s, text}, space)
		s = strings.Join([]string{s, "hundred"}, " ")
		space = " "
		num = num % 100
	}
	if num >= 20 && num < 100 {
		text, _ := GetDecimalPart(int(math.Floor(float64(num) / 10)))
		if len(s) > 0 {
			s = strings.Join([]string{s, text}, " and ")
		} else {
			s = strings.Join([]string{s, text}, space)
		}
		num = num % 10
		text2, _ := GetNumberTextUpToTen(num)
		s = strings.Join([]string{s, text2}, "-")
		num = 0
	}
	if num > 0 {
		text, _ := GetNumberTextUpToTen(num)
		if len(s) > 0 {
			s = strings.Join([]string{s, text}, " and ")
		} else {
			s = strings.Join([]string{s, text}, "")
		}

	}
	return s
}
func GetDecimalPart(num int) (string, error) {
	switch num {
	case 2:
		return "twenty", nil
	case 3:
		return "thirty", nil
	case 4:
		return "forty", nil
	case 5:
		return "fifty", nil
	case 6:
		return "sixty", nil
	case 7:
		return "seventy", nil
	case 8:
		return "eighty", nil
	case 9:
		return "ninety", nil
	}
	return "", errors.New("Invalid number")
}

func GetNumberTextUpToTen(num int) (string, error) {
	switch num {
	case 1:
		return "one", nil
	case 2:
		return "two", nil
	case 3:
		return "three", nil
	case 4:
		return "four", nil
	case 5:
		return "five", nil
	case 6:
		return "six", nil
	case 7:
		return "seven", nil
	case 8:
		return "eight", nil
	case 9:
		return "nine", nil
	case 10:
		return "ten", nil
	case 11:
		return "eleven", nil
	case 12:
		return "twelve", nil
	case 13:
		return "thirteen", nil
	case 14:
		return "fourteen", nil
	case 15:
		return "fifteen", nil
	case 16:
		return "sixteen", nil
	case 17:
		return "seventeen", nil
	case 18:
		return "eighteen", nil
	case 19:
		return "nineteen", nil
	}
	return "", errors.New("Invalid number")

}
