package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() (command, args string) {
	buf := bufio.NewReader(os.Stdin)
	argsArray:= []string{}
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		argsArray = strings.Split(string(sentence), " ")
	}
	//remove any empty strings at beginning
	for len(argsArray) > 0 && argsArray[0] == "" {
		argsArray = argsArray[1:]
	}
	if len(argsArray) < 1 {
		fmt.Println("No command found")
		return
	}
	command = strings.Trim(strings.ToUpper(argsArray[0]), "\n")
	args = strings.Trim(strings.Join(argsArray[1:], " "), "\n")
	fmt.Printf("I read; \"%s %s\" \n", command, args)

	return
}
