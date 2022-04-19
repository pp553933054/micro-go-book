package main

import (
	"fmt"
	"github.com/pp553933054/micro-go-book/ch4-feature/compute"
)

func main() {

	params := &compute.IntParams{
		P1: 1,
		P2: 2,
	}
	fmt.Println(params.Add())

}
