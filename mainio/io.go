package mainio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

func GetLocation() string {
	usr, _ := user.Current()
	return fmt.Sprintf("%s/.holdthisdb", usr.HomeDir)
}

func FileCheck(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Making file", filename)
		ioutil.WriteFile(filename, []byte("[{\"key\":\"roothome\",\"value\":\"/root\"}]"), 0664)
		return
	}
}

type Entry struct {
	Key   string
	Value string
}

func GetEntries() []Entry {
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
	panic("Can't continue")
}

func Get(key string) Entry {
	Bits := GetEntries()
	for _, testcase := range Bits {
		if testcase.Key == key {
			return testcase
		}
	}
	panic("Cannot Find")
}

func Set(key string, value string) {
	Bits := GetEntries()
	e := Entry{
		Key:   key,
		Value: value,
	}
	// fmt.Println(e)
	BitsReloaded := append(Bits, e)
	b, _ := json.Marshal(BitsReloaded)
	ioutil.WriteFile(GetLocation(), b, 0664)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
