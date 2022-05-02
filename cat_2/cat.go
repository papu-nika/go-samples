package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	argc := len(os.Args)
	if argc == 1 {
		if err := input_and_write(os.Stdin.Name()); err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	} else {
		flag.Parse()
		filenames := flag.Args()
		for i := 0; argc-1 > i; i++ {
			if err := input_and_write(filenames[i]); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func input_and_write(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for {
		data := make([]byte, 1024)
		count, err := f.Read(data)
		if count == 0 {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s", string(data[:count]))
	}
}
