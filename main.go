package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var tabB [][]string

func main() {

	if len(os.Args) != 2 {
		os.Exit(1)
	}
	argu := os.Args[1]
	data, _ := ioutil.ReadFile("standard.txt")
	linsplit := string(data)
	str := strings.Split(linsplit, "\n")
	var tabA []string
	for _, j := range str {
		if len(j) != 0 {

			tabA = append(tabA, string(j))
		} else {

			tabB = append(tabB, tabA)
			tabA = []string{}
		}

	}
	printAscii(argu)

}
func printAscii(str string) {
	var i int

	s := strings.Split(str, "\\n")

	for _, j := range s {

		for i < 8 {
			for _, elem := range j {
				fmt.Print(tabB[byte(elem)-31][i], " ")

			}
			i++
			fmt.Println()

		}
		i = 0

	}

}
