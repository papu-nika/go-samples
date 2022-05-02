package main

import (
	"fmt"
	"os"
	"strconv"
)

type marubatu [][]string

type pattern [][]string

func main() {
	mb := marubatu{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	for {
		mb.Print()
		fmt.Printf("プレイヤーA(o) : 数値を入力してください\n> ")
		mb.Write("o")
		if mb.Check() {
			fmt.Println("\nAの勝利")
			mb.Print()
			os.Exit(0)
		}
		mb.Print()
		fmt.Printf("プレイヤーB(x) : 数値を入力してください\n> ")
		mb.Write("x")
		if mb.Check() {
			fmt.Println("\nBの勝利")
			mb.Print()
			os.Exit(0)
		}
	}
}

func (mb marubatu) Print() {
	i := 1
	fmt.Printf("\n")
	for _, y := range mb {
		for _, x := range y {
			fmt.Printf("%s ", x)
		}
		fmt.Printf(" |  %d %d %d\n", i, i+1, i+2)
		i += 3
	}
}

func (mb marubatu) Write(symbol string) bool {

	for {
		var input string
		fmt.Scan(&input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("1~9の数字を入力してください\n> ")
			continue
		}
		if !(0 < num && num < 10) {
			fmt.Printf("1~9の数字を入力してください\n> ")
			continue
		} else {
			n := mb.Calc(num)
			if *n != "-" {
				fmt.Printf("既に埋まっています。他の数字を選択してください\n> ")
				continue
			} else {
				*n = symbol
				break
			}
		}
	}

	return true
}

func (mb marubatu) Calc(num int) *string {
	y := (num - 1) / 3
	x := (num - 1) % 3
	return &mb[y][x]
}

func (mb marubatu) Check() bool {

	for num := 1; num < 10; num += 3 {
		if result := mb.Check_Rank(num); result == true {
			return result
		}
	}
	for num := 1; num < 4; num++ {
		if result := mb.Check_Column(num); result == true {
			return result
		}
	}
	if result := mb.Check_Parastichy(); result == true {
		return result
	}

	return false
}

func (mb marubatu) Check_Rank(num int) bool {
	if *(mb.Calc(num)) == *(mb.Calc(num + 1)) &&
		*(mb.Calc(num)) == *(mb.Calc(num + 2)) &&
		*(mb.Calc(num)) != "-" {
		return true
	}
	return false
}

func (mb marubatu) Check_Column(num int) bool {
	if *(mb.Calc(num)) == *(mb.Calc(num + 3)) &&
		*(mb.Calc(num)) == *(mb.Calc(num + 6)) &&
		*(mb.Calc(num)) != "-" {
		return true
	}
	return false
}

func (mb marubatu) Check_Parastichy() bool {
	if *(mb.Calc(1)) == *(mb.Calc(5)) &&
		*(mb.Calc(1)) == *(mb.Calc(9)) &&
		*(mb.Calc(1)) != "-" {
		return true
	} else if *(mb.Calc(3)) == *(mb.Calc(5)) &&
		*(mb.Calc(3)) == *(mb.Calc(7)) &&
		*(mb.Calc(3)) != "-" {
		return true
	}
	return false
}
