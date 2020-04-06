package main

import (
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

func (this BranchingReport) _getMaxLen() int {
	maxLen := 0
	for _, arr := range this.Lines {
		for _, str := range arr {
			if len(str) > maxLen {
				maxLen = len(str)
			}
		}
	}
	return maxLen
}

func (this *BranchingReport) addOR() {
	this.Lines = append(this.Lines, []string{ORSIGNAL})
}

func (this *BranchingReport) initialize() {
	this.Lines = [][]string{[]string{}}
}

func (this BranchingReport) String() (out string) {
	for _, line := range this.Lines {
		out = out + "\n" + strings.Join(line, " ")
	}
	return
}
