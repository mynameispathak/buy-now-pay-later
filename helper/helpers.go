package helper

import (
	"bufio"
	"fmt"
	"strings"
)

// Reads space separated values from reader and splits them into a list
func ReadVals(reader *bufio.Reader) []string {
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Something bad happened")
		panic(err)
	}

	text = strings.TrimSuffix(text, "\n")
	vals := strings.Split(text, " ")
	return vals
}
