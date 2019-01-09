package main

import "testing"

func TestTask17(t *testing.T) {
	oneThousand := Task17(1000)
	if oneThousand != "one thousand" {
		t.Error("Expected one thousand, got", oneThousand)
	}
	threeThousand := Task17(3009)
	if threeThousand != "three thousand and nine" {
		t.Error("Expected three thousand and nine, got", threeThousand)
	}
}
