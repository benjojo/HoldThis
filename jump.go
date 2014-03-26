package main

import (
	mainio "./mainio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 2 {

		target := mainio.Get(os.Args[1]).Value
		os.Chdir(target)
		os.Setenv("PWD", target)
		fmt.Println("cd", os.Getenv("PWD"))
	} else {
		// fmt.Println("HoldThis - Jump Program")
		// fmt.Println("$ jump <name of hold>")
		os.Exit(1)
	}
}
