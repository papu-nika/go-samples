package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	arg := flag.Args()
	fmt.Println(arg)
	f, err := os.Open(os.Stdin.Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log_file := log_init()
	defer log_file.Close()

	buf, i := read_file(f)
	//QuicSort(buf, 0, len(buf))
	for v := 0; v < i; v++ {
		fmt.Println(buf[v])
	}
	
	sort.Strings(buf)

	for v := 0; v < i; v++ {
		fmt.Println(buf[v])
	}
}

func read_file(f *os.File) ([]string, int) {
	now := time.Now()
	//var buf []string
	buf := make([]string, 10000000)
	read := bufio.NewScanner(f)

	var i int
	for i = 1; read.Scan(); i++ {
		buf[i] = read.Text()
	}
	log.Printf("##Read_File##\t%d milisecond\tkey= \"%s\"", time.Since(now).Milliseconds(), "")
	return buf, i
}

func QuicSort(buf *[10000000]string, start, last int) {
	now := time.Now()
	var qart int
	if last-start > 15 {
		InsertSort(buf, start, last)
	} else if start < last {
		qart = Quicsort_part_left_right(buf, start, last)
		go func(s, l int) {
			QuicSort(buf, s, l)
		}(start, qart-1)
		go func(s, l int) {
			QuicSort(buf, s, l)
		}(qart+1, last)
	}
	log.Printf("##Sort##\t%d milisecond\tkey= \"%s\"", time.Since(now).Milliseconds(), "")
	return
}

func Quicsort_part_left_right(buf *[10000000]string, start, last int) int {
	i := start - 1
	pivot := (*buf)[last]
	for k := start; k < last; k++ {
		if (*buf)[k] < pivot {
			i++
			Swap_index_buf(buf, i, k)
		}
	}
	Swap_index_buf(buf, i+1, last)
	return i + 1
}

func InsertSort(buf *[10000000]string, start, last int) {

	for ; start < last; start++ {
		for i := 0; i < start; i++ {
			if (*buf)[start-i-1] > (*buf)[start-i] {
				Swap_index_buf(buf, start-i-1, start-i)
			} else {
				break
			}
		}
	}
}

func Swap_index_buf(buf *[10000000]string, a, b int) {
	(*buf)[a], (*buf)[b] = (*buf)[b], (*buf)[a]
}

func log_init() *os.File {
	file, err := os.OpenFile("_test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		is_error(err)
	}
	log.SetOutput(file)
	return file
}

func is_error(err error) {
	fmt.Println(err)
	os.Exit(1)
}
