package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	fmt.Println(Task17(arg))
}
