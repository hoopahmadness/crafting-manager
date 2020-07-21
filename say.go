package main

import (
	"fmt"
	// "github.com/aybabtme/rgbterm"
	"gopkg.in/gookit/color.v1"
)

type Saying struct {
	NONE      func(...interface{}) string
	HELP      func(...interface{}) string
	DEFAULT   func(...interface{}) string
	RESULT    func(...interface{}) string
	DIRECTION func(...interface{}) string
	QUESTION  func(...interface{}) string
	EMPHASIS  func(...interface{}) string
	WARNING   func(...interface{}) string
}

func renderColors() Saying {
	say := Saying{}
	say.DEFAULT = color.FgGreen.Render
	say.HELP = color.FgWhite.Render
	say.RESULT = color.FgMagenta.Render
	say.DIRECTION = color.FgCyan.Render
	say.QUESTION = color.FgBlue.Render
	say.EMPHASIS = color.Bold.Render
	say.WARNING = color.FgYellow.Render
	say.NONE = func(obj ...interface{}) string {
		return fmt.Sprintf("%v", obj)
	}

	return say
}

func (this Saying) paintWord(str string, X, Y int) string {
	//for now...
	// return say.WARNING(str)
	return str
}
