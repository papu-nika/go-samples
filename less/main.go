package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault

var search_str string

var search_str_index int

type File_buf []struct {
	line string
}

func main() {
	b := make(File_buf, 100000)
	buf := &b

	file := log_init()
	defer file.Close()
	flag.Parse()
	argments := flag.Args()

	if err := termbox.Init(); err != nil {
		is_error(err)
	}
	defer termbox.Close()
	buf.ReadFile(argments[0])
	log.Print(cap(*buf), len(*buf))

	buf.DrawBox(0, "")
MAINLOOP:
	for i := 0; ; {
		//log.Printf("%p", &search_str)
		//w, h := termbox.Size()
		// (2) イベントハンドリング
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break MAINLOOP
			case termbox.KeyArrowRight, termbox.KeyArrowLeft, termbox.KeyArrowUp, termbox.KeyArrowDown:
				buf.Keyarrow_procces(&i, ev.Key)
				continue MAINLOOP
			case termbox.KeyBackspace2:
				search_str = search_str[:len(search_str)-1]
				buf.DrawBox(i, search_str)
				continue MAINLOOP
			}
		}
		search_str += string(ev.Ch)
		buf.DrawBox(i, search_str)
	}
}

func log_init() *os.File {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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

func (buf *File_buf) Keyarrow_procces(i *int, ev termbox.Key) {
	switch ev {
	case termbox.KeyArrowDown:
		*i++
		buf.DrawBox(*i, search_str)
	case termbox.KeyArrowUp:
		if *i == 0 {
			return
		} else {
			*i--
			buf.DrawBox(*i, search_str)
		}
	}
	return
}

func (buf *File_buf) ReadFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	read := bufio.NewReader(f)

	var i uint64
	for i = 0; ; i++ {
		line, err := read.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		if i < uint64(cap(*buf)) {
			test1 := make(File_buf, 1, 1)
			*buf = append((*buf), test1[0])

		}
		(*buf)[i].line = string(line)
	}
	log.Print(cap(*buf))
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		if c == '\t' {
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
		} else {
			termbox.SetCell(x, y, c, fg, bg)
			x += runewidth.RuneWidth(c)
		}
	}
}

func (buf *File_buf) DrawBox(start_index int, input string) {
	index := start_index
	_, h := termbox.Size()
	termbox.Clear(coldef, coldef)
	tbprint(0, 0, coldef, coldef, input)
	if search_str == "" {
		for i := 1; i < h; i++ {
			tbprint(0, i, coldef, coldef, (*buf)[index].line)
			index++
		}
	} else {
		for i := 1; i < h; {
			s, err := regexp.Compile(search_str)
			if err != nil {
				os.Exit(0)
			}
			b := s.Find([]byte((*buf)[index].line))
			if b == nil {
				index++
				continue
			} else {
				tbprint(0, i, coldef, coldef, (*buf)[index].line)
				i++
				index++
			}
		}
	}

	termbox.Flush() // Flushを呼ばないと描画されない
}
