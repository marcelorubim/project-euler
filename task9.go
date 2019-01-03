package main

func Task9() int {
	product := 0
	sum := 0
	lim := 1000
	for a := 0; a < lim; a++ {
		for b := 0; b < lim; b++ {
			for c := 0; c < lim; c++ {
				a2 := a * a
				b2 := b * b
				c2 := c * c
				if a < b && b < c && (a2+b2) == c2 {
					product = a * b * c
					sum = a + b + c
					if sum == 1000 {
						return product
					}
				}
			}
		}
	}

	return product
}
