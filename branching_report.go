package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ORSIGNAL = "OR"
)

type BranchingReport struct {
	Lines [][]string
}

// sticks a string at the end of the array at the end of the arrays
//panics
func (this *BranchingReport) appendString(str string) {
	if len(this.Lines) == 0 {
		this.initialize()
	}

	ind := len(this.Lines) - 1
	this.Lines[ind] = append(this.Lines[ind], str)
}

//inserts a string in front of every []string in this array of arrays
func (this *BranchingReport) insertString(str string) {
	for ind, line := range this.Lines {
		isOR := false
		if len(line) > 0 {
			if line[len(line)-1] == ORSIGNAL {
				isOR = true
			}
		}
		if !isOR {
			this.Lines[ind] = append([]string{str}, line...)
		} else {
			this.Lines[ind] = append([]string{strings.Repeat(" ", len(str)+5)}, line...)

		}
	}
}

func (this *BranchingReport) combineReports(other BranchingReport) {
	newLines := append(this.Lines, other.Lines...)
	this.Lines = newLines
}

// func (this BranchingReport) _getMaxLen() int {
// 	maxLen := 0
// 	for _, arr := range this.Lines {
// 		for _, str := range arr {
// 			if len(str) > maxLen {
// 				maxLen = len(str)
// 			}
// 		}
// 	}
// 	return maxLen
// }

func (this BranchingReport) getMaxLevels() int {
	lev := 0
	for _, line := range this.Lines {
		if len(line) > lev {
			lev = len(line)
		}
	}
	return lev
}

func (this *BranchingReport) addOR() {
	this.Lines = append(this.Lines, []string{ORSIGNAL})
}

func (this *BranchingReport) initialize() {
	this.Lines = [][]string{[]string{}}
}

func (this BranchingReport) String() (out string) {
	this.processLines()
	for _, line := range this.Lines {
		out = out + "\n" + strings.Join(line, "*")
	}
	return
}

func (this *BranchingReport) processLines() {
	levels := this.getMaxLevels()
	Ysections := splitRange(defaultColorRange[3], defaultColorRange[2], levels)
	for indY, line := range this.Lines {
		for indX, phrase := range line {
			//Craft #-80#^1 Super Net^ from #-80#^1 Gold Net^;
			if len(strings.Trim(phrase, " ")) != 0 && phrase != ORSIGNAL {
				tokens := strings.Split(phrase, "#")
				report("", "", tokens)
				col1, err := strconv.Atoi(tokens[1])
				if err != nil {
					fmt.Println(err)
				}
				col2, err := strconv.Atoi(tokens[3])
				if err != nil {
					fmt.Println(err)
				}

				middleSection := strings.Split(tokens[2], "^")
				lastSection := strings.Split(tokens[4], "^")
				word1 := say.paintWord(middleSection[1], col1, Ysections[indX])
				word2 := say.paintWord(lastSection[1], col2, Ysections[indX])

				newLine := tokens[0] + middleSection[0] + word1 + middleSection[2] + lastSection[0] + word2 + lastSection[2]
				fmt.Println(newLine)
				this.Lines[indY][indX] = newLine
			}
		}
	}
}
