package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func GetLocation() string {
	usr, _ := user.Current()
	return fmt.Sprintf("%s/.holdthisdb", usr.HomeDir)
}

func FileCheck(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Making file %s\n", filename)
		Sample := make([]Entry, 0)
		SE := Entry{"roothome", "/root"}
		Sample = append(Sample, SE)
		b, _ := json.Marshal(&Sample)
		ioutil.WriteFile(filename, b, 0664)
		return
	}
}

type Entry struct {
	Key   string
	Value string
}

func GetEntries() (ret []Entry) {
	path := GetLocation()
	FileCheck(path)
	file, e := ioutil.ReadFile(path)
	check(e)
	var Bits []Entry
	err := json.Unmarshal(file, &Bits)
	// fmt.Println(Bits)
	if err == nil {
		return Bits
	} else {
		fmt.Println("Oh dear. I can't decode our bookmarks file. this is a pretty big issue.")
	}
	log.Fatal("Can't continue")
	return ret
}

func Get(key string) (e Entry) {
	Bits := GetEntries()
	for _, testcase := range Bits {
		if testcase.Key == key {
			return testcase
		}
	}
	log.Fatal("Cannot Find")
	return e
}

func Set(key string, value string) {
	Bits := GetEntries()
	e := Entry{
		Key:   key,
		Value: value,
	}
	// fmt.Println(e)
	BitsReloaded := append(Bits, e)
	write(BitsReloaded)
}

func Remove(key string) {
	Bits := GetEntries()
	for i, testcase := range Bits {
		if testcase.Key == key {
			Bits = append(Bits[:i], Bits[i+1:]...)
			fmt.Println("Removed the directory", testcase.Value, "with the name '", testcase.Key, "'")
		}
	}
	write(Bits)
}

func write(Bits []Entry) {
	b, _ := json.Marshal(Bits)
	ioutil.WriteFile(GetLocation(), b, 0664)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
