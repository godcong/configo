package main

import "fmt"

func main() {
	p := config.Get(`default`).Get("path")
	fmt.Println(p)

}
