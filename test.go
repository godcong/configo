package main

import (
	"fmt"

	"github.com/godcong/configo/config"
)

func main() {
	config.Load()
	p := config.Get(`default`).Get("path")
	fmt.Println(p)

}
