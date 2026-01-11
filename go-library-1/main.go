package main

import (
	"fmt"

	"github.com/AlexDiwong/build-testing/go-library-1/arithmetic"
)

func main() {
	fmt.Println("hello, twice")
	res := arithmetic.Add(1,3)
	fmt.Println(res)
}
