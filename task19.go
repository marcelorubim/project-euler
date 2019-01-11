package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Date struct {
	day       int
	month     int
	year      int
	dayOfWeek int
}

func (d *Date) Text() string {
	return strings.Join([]string{strconv.Itoa(d.day), strconv.Itoa(d.month), strconv.Itoa(d.year)}, "/")
}

func Task19() int {
	count := 0

	date := Date{1, 1, 1900, 2}

	for date.Text() != "31/12/2000" {
		if date.year >= 1901 && date.dayOfWeek == 1 && date.day == 1 {
			count++
		}
		fmt.Println(date)
		date = calculateNextDay(date)

	}

	return count
}

func calculateNextDay(date Date) Date {
	limitDay := 31
	if date.month == 4 || date.month == 6 || date.month == 9 || date.month == 11 {
		limitDay = 30
	} else if date.month == 2 {
		if isLeapYear(date.year) {
			limitDay = 29
		} else {
			limitDay = 28
		}
	}
	if date.day+1 > limitDay {
		date.day = 1
		if date.month == 12 {
			date.month = 1
			date.year++
		} else {
			date.month++
		}
	} else {
		date.day++
	}
	if date.dayOfWeek == 7 {
		date.dayOfWeek = 1
	} else {
		date.dayOfWeek++
	}
	return date
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
