package main

import (
	"fmt"
	"os"
)

func swap(one, two string) (err error) {
	tmp := ".tmp." + one
	if err = os.Rename(one, tmp); err != nil {
		return
	}
	if err = os.Rename(two, one); err != nil {
		// try to cleanup
		defer os.Rename(tmp, one)
		return
	}
	if err = os.Rename(tmp, two); err != nil {
		// try to cleanup
		defer os.Rename(tmp, one)
		defer os.Rename(one, two)
		return
	}
	return
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("must provide exactly two files to swap")
		return
	} else if args[0] == args[1] {
		// just do nothing if they are the same file
		return
	}
	err := swap(args[0], args[1])
	if err != nil {
		fmt.Println(err)
	}
}
