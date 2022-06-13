package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	err := exec.Command("su", "test").Start()
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("id").Output()
	fmt.Println(string(out))
}
