package main

import (
	"fmt"

	"github.com/psinthorn/gomodules/toolkit"
)

func main() {
	var tools toolkit.Tools
	s := tools.RandomString(10)
	fmt.Println(s)
}
