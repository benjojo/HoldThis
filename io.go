package main

import (
	"fmt"
	"os"
	"os/user"
)

func GetLocation() string {
	usr, err := user.Current()
	if err != nil {
		return fmt.Sprintf("%s/.holdthisdb", usr)
	} else {
		panic(err)
	}
}

func Get() {
	path := GetLocation()
}

func Set() {

}
