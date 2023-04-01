package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var tabB [][]string

func main() {

	if len(os.Args) != 4 {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	outputfile := os.Args[1]
	check := strings.Split(outputfile, "=")
	if len(check) != 2 || check[0] != "--output" {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)

	}
	argu := os.Args[2]
	banner := os.Args[3]

	banner = banner + ".txt"
	data, err := ioutil.ReadFile(banner)
	if err != nil {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
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
	errr := ioutil.WriteFile(check[1], []byte(printAscii(argu)), 0644)
	if err != nil {
		log.Fatal(errr)
	}

}
func printAscii(str string) string {
	var i int
	outpout := ""

	s := strings.Split(str, "\\n")

	for _, j := range s {

		if j != "" && len(j) == 1 && []byte(j)[0]-0 < 32 || []byte(j)[0] > 126-0 {
			outpout = ""

		} else {
			for i < 8 {
				if len(j) == 0 {
					outpout += ""

				}

				for _, elem := range j {
					if len(string(elem)) == 0 {
						outpout += ""

					}
					if byte(elem) >= 32 && byte(elem) <= 126 {

						outpout += tabB[byte(elem)-31][i] + " "
					}

					//fmt.Print(tabB[byte(elem)-31][i], " ")

				}
				i++
				outpout = outpout + "\n"

			}
			i = 0

		}

	}
	return outpout

}
