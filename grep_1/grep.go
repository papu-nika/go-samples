package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var err error
	argc := len(os.Args)
	flag.Parse()
	argments := flag.Args()
	if argc == 1 {
		fmt.Println("Few argments")
	} else if argc == 2 {
		err = do_grep(os.Stdin.Name(), argments[0])
	} else if argc == 3 {
		err = do_grep(argments[1], argments[0])
	} else {
		fmt.Println("Too many argments")
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func do_grep(file string, str string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	read := bufio.NewReader(f)
	s, err := regexp.CompilePOSIX(str)
	for i := 0; ; i++ {
		line, err := read.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		b := s.FindStringIndex(line)
		if b == nil {
			continue
		} else {
			fmt.Printf("%s", line)
		}
	}
	return nil
}
