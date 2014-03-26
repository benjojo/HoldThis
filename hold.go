package main

import (
	mainio "./mainio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 2 {
		str, _ := os.Getwd()
		fmt.Println("Saving the directory", str, "as the name '", os.Args[1], "'")
		mainio.Set(os.Args[1], str)
	} else {
		fmt.Println("HoldThis - Hold program")
		fmt.Println("$ hold <name of hold>")
		fmt.Println("you can then jump to the hold by doing")
		fmt.Println("$ jump <name of hold>")
		os.Exit(1)
	}
}
