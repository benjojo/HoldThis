package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(len(os.Args))
	List := GetEntries()
	for _, v := range List {
		if len(os.Args) > 1 {
			if strings.HasPrefix(v.Key, os.Args[1]) {
				fmt.Printf("%s -> %s\n", v.Key, v.Value)
			}
		} else {
			fmt.Printf("%s -> %s\n", v.Key, v.Value)
		}
	}
}
