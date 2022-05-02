package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	argments := flag.Args()
	argc := len(argments)

	if argc == 0 {
		do_find(".")
	} else {
		for i := 0; i < argc; i++ {
			do_find(argments[i])
		}
	}
}

func do_find(f string) {
	file_info, err := os.Stat(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	if file_info.IsDir() == false {
		return
	}
	fmt.Println(f)
	files, err := ioutil.ReadDir(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(files); i++ {
		file := filepath.Join(f, files[i].Name())
		fmt.Println(file)
		do_find(filepath.Join(f, files[i].Name()))
	}
	return
}
