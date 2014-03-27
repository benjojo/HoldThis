package main

import (
	"fmt"
)

func main() {
	List := GetEntries()
	for _, v := range List {
		fmt.Printf("%s -> %s\n", v.Key, v.Value)
	}
}
