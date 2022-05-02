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

type File_buf struct {
	line *string
	node *File_buf
}

func main() {
	var buf File_buf

	file := log_init()
	defer file.Close()
	flag.Parse()
	argments := flag.Args()

	if err := termbox.Init(); err != nil {
		is_error(err)
	}
	defer termbox.Close()
	buf.ReadFile(argments[0])

	buf.DrawBox(0, 0, "")

	high, width := 0, 0
MAINLOOP:
	for {
		//log.Printf("%p", &search_str)
		// (2) イベントハンドリング
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break MAINLOOP
			case termbox.KeyArrowRight, termbox.KeyArrowLeft, termbox.KeyArrowUp, termbox.KeyArrowDown:
				buf.Keyarrow_procces(&high, &width, ev.Key)
				continue MAINLOOP
			case termbox.KeyBackspace2:
				if len(search_str) == 0 {
					continue MAINLOOP
				}
				search_str = search_str[:len(search_str)-1]
				buf.DrawBox(high, width, search_str)
				continue MAINLOOP
			}
		}
		search_str += string(ev.Ch)
		buf.DrawBox(high, width, search_str)
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

func (buf *File_buf) Keyarrow_procces(high *int, width *int, ev termbox.Key) {
	switch ev {
	case termbox.KeyArrowDown:
		*high++
		buf.DrawBox(*high, *width, search_str)
	case termbox.KeyArrowUp:
		if *high == 0 {
			return
		} else {
			*high--
			buf.DrawBox(*high, *width, search_str)
		}
	case termbox.KeyArrowRight:
		*width++
		buf.DrawBox(*high, *width, search_str)
	case termbox.KeyArrowLeft:
		if *width == 0 {
			return
		}
		*width--
		buf.DrawBox(*high, *width, search_str)
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

	for {
		line, err := read.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		var new_buf File_buf
		buf.line = &line
		buf.node = &new_buf
		buf = &new_buf
	}
	buf.node = nil
}

func (buf *File_buf) DrawBox(high int, width int, input string) {
	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()
	log.Print(w)
	tbprint(0, 0, coldef, coldef, input)

	if search_str == "" {
		for i := 0; i < high; i++ {
			buf = buf.node
		}
		for i := 1; i < h; i++ {
			tbprint(w*width, i, coldef, coldef, *(buf.line))
			buf = buf.node
		}
	} else {
		line := 0
		for i := 1; i < h && (*buf).node != nil; {
			// s, err := regexp.MustCompile(search_str)
			// if err != nil {
			// 	os.Exit(0)
			// }
			// b := s.Find([]byte(buf.line))
			b, e := regexp.MatchString(search_str, *(buf.line))
			if e != nil {
				os.Exit(0)
			}
			if b == false {
				buf = buf.node
				continue
			} else {
				if line < high {
					buf = buf.node
					line++
				} else {
					tbprint(w*width, i, coldef, coldef, *(buf.line))
					buf = buf.node
					i++
				}
			}
		}
	}

	termbox.Flush() // Flushを呼ばないと描画されない
}

func tbprint(high, y int, fg, bg termbox.Attribute, str string) {
	x := 0
	for i := 0; i < len(str); i++ {
		if i < high {
			continue
		}
		if str[i] == '\t' {
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
		} else {
			termbox.SetCell(x, y, rune(str[i]), fg, bg)
			x += runewidth.RuneWidth(rune(str[i]))
		}
	}
}
