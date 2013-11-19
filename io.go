package main

import (
	"fmt"
	"io/ioutil"
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

func FileCheck(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Making file", filename)
		ioutil.WriteFile(filename, []byte("[]"), 0664)
		return
	}
}

func Get() {
	path := GetLocation()
	FileCheck(path)
}

func Set() {
	path := GetLocation()
	FileCheck(path)

}
