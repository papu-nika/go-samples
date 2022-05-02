package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "test.txt"
	f, err := os.Open(filename)
	if err != nil {
		write_hello(filename)
		os.Exit(0)
	}
	defer f.Close()
	for {
		data := make([]byte, 1024)
		count, err := f.Read(data)
		if count == 0 {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s", string(data[:count]))
	}
}

func write_hello(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	str := "Hello World"
	data := []byte(str)
	count, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("write %d bytes\n", count)
}
