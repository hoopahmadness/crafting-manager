package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func parseInput() string {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		say(err)
	}
	return strings.Trim(string(sentence), "\n")
}

func getCommandAndArgs() (command, args string) {
	sentence := parseInput()
	argsArray := strings.Split(string(sentence), " ")
	//remove any empty strings at beginning
	for len(argsArray) > 0 && argsArray[0] == "" {
		argsArray = argsArray[1:]
	}
	if len(argsArray) < 1 {
		return
	}
	command = strings.ToUpper(argsArray[0])
	args = strings.Join(argsArray[1:], " ")

	return

}

func getYesOrNo() (answer string) {
	for answer != YES && answer != NO {
		say("Please send YES or NO ")
		answer = strings.ToUpper(parseInput())	
	}
	return
}

func getBackResponse() (answer bool) {
	answer = strings.ToUpper(parseInput())==BACK
	return
}

func getInteger() (answer int) {
	noGoodAnswer := true
	var err error
	for noGoodAnswer {
		say("Please send a valid whole number")
		answerStr := parseInput()
		answer, err = strconv.Atoi(answerStr)
		if err == nil {
			noGoodAnswer = false
		} else {
			say(err)
		}
	}
	return
}