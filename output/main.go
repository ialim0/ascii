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
	var i, n int
	outpout := ""
	if len(str) == 0 {
		return ""
	}

	s := strings.Split(str, "\\n")

	for _, j := range s {
		if j == "" {
			outpout = outpout + "\n"
		} else {
			for i < 8 {

				for _, elem := range j {
					if byte(elem) >= 32 && byte(elem) <= 126 {
						n = 1

						outpout += tabB[byte(elem)-31][i] + ""
					}

				}

				outpout = outpout + "\n"

				i++

			}
			i = 0

		}

	}

	if n == 1 {
		return outpout
	}
	outpout = strings.Replace(outpout, "\n", "", 1)
	return outpout
}
