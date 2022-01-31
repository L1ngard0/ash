package main

import (
	"fmt"

	"github.com/L1ngard0/ash/blockchain"
)

func main() {
	var x map[string]int
	x["key"] = 10
	fmt.Print(x)
	fmt.Print("qwe")
	blockchain.Test()
}
