package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var d = flag.String("d", "", "delete entry")

func main() {
	flag.Parse()

	if *d != "" {
		Remove(*d)
		os.Exit(0)
	}

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
