package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input = bufio.NewScanner(os.Stdin)

func main() {
	str := gets()
	char := gets()
	if strings.Contains(str, char) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func gets() string {
	if input.Scan() {
		return input.Text()
	} else {
		return "NoInputData"
	}
}
