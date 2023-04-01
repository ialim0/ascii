package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var tabB [][]string

func main() {

	if len(os.Args) != 3 {
		fmt.Print("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		os.Exit(0)
	}
	argu := os.Args[1]
	banner := os.Args[2]

	banner = banner + ".txt"
	data, err := ioutil.ReadFile(banner)
	if err != nil {
		fmt.Println("Unable to open file")
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

	fmt.Print(printAscii(argu))

}
func printAscii(str string) string {
	var i int
	outpout := ""

	s := strings.Split(str, "\\n")

	for k, j := range s {
		if j==""{
			if k<len(s)-1{
				outpout=outpout+"\n"

			}
			
		}else{
			for i < 8 {

				for _, elem := range j {
					if byte(elem) >= 32 && byte(elem) <= 126 {

						outpout += tabB[byte(elem)-31][i] + " "
					}

					//fmt.Print(tabB[byte(elem)-31][i], " ")

				}
				i++
				if i<7 && j!=""{
					outpout = outpout + "\n"

				} 
				

			}
			i = 0

		}

	
		

		

	}
	return outpout

}
