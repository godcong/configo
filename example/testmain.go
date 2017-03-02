package main

import (
	"fmt"

	"github.com/godcong/configo"
)

func main() {
	p := configo.Get(`default`).Get("path")
	fmt.Println(p)

}
