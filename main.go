package main

import (
	"fmt"

	"github.com/marcsantiago/go-ua-benchmarks/usurf"
	"github.com/marcsantiago/go-ua-benchmarks/xojoc"
)

func main() {
	isValid := xojoc.IsValid(usurf.THBadUserAgent)
	fmt.Println(isValid)
}
