package mainio

import (
	"encoding/json"
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

type Entry struct {
	key   string
	value string
}

func GetEntries() []Entry {
	path := GetLocation()
	FileCheck(path)
	file, e := ioutil.ReadFile(path)
	check(e)
	var Bits []Entry
	err := json.Unmarshal(file, &Bits)
	if err != nil {
		return Bits
	} else {
		fmt.Println("Oh dear. I can't decode our bookmarks file. this is a pretty big issue.")
	}
	panic("Can't continue")
}

func Get(key string) Entry {
	Bits := GetEntries()
	for _, testcase := range Bits {
		if testcase.key == key {
			return testcase
		}
	}
	panic("Cannot Find")
}

func Set(key string, value string) {
	Bits := GetEntries()
	var e Entry
	e.key = key
	e.value = value
	BitsReloaded := append(Bits, e)
	b, _ := json.Marshal(BitsReloaded)
	ioutil.WriteFile(GetLocation(), b, 0664)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
