package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func testingg(t *testing.T) {
	var newTab, resultTab []string
	data, _ := ioutil.ReadFile("input_test.txt")
	newTab = strings.Split(string(data), "\n")
	for i := 0; i < len(newTab); i++ {
		resultTab = append(resultTab, printAscii(newTab[i]))

	}

	//tabResult := strings.Join(resultTab, "\n")
	data2, _ := ioutil.ReadFile("expected.txt")
	exp := strings.Split(string(data2), "\n")

	if len(exp) != len(resultTab) {
		t.Errorf("PrintAscii() = %q, expected %q", resultTab, exp)

	} else {
		for i := 0; i < len(resultTab); i++ {
			if resultTab[i] != exp[i] {
				t.Errorf("PrintAscii() = %q, expected %q", resultTab, exp)

			}
		}

	}

}
