package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	current_dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	files, err := ioutil.ReadDir(current_dir)
	var i int
	for i = 0; len(files)-1 > i; i++ {
		fmt.Printf("%s  ", files[i].Name())
	}
	fmt.Println(files[i].Name())
}
