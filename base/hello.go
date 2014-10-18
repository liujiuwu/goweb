package main

import (
	"errors"
	"fmt"
)

var vname1, vname2, vname3 = "v1", "v2", "v3"

const (
	book = "go lang"
)

func main() {
	name := "liujiuwu"
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
	fmt.Println(name + "|" + book)

	cname := []byte(name)
	cname[0] = 'w'

	fmt.Println(string(cname))

	mt := `test
	       name
	       liujiuwu`

	fmt.Println(mt)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
