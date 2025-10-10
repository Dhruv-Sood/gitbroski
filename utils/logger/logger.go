// Package logger provides colorized console output utilities.
package logger

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func Banner(text string) {
	myFigure := figure.NewFigure(text, "slant", true)
	c := color.New(color.FgHiMagenta)
	_, _ = c.Println(myFigure.String())
}

func Text(text string) {
	c := color.New()
	c.AddRGB(135, 206, 250)
	_, _ = c.Println(text)
}

func Error(text string) {
	c := color.New(color.FgHiRed)
	_, _ = c.Println(text)
}

func Success(text string) {
	c := color.New(color.FgHiGreen)
	c.Println(text)
}

func Warning(text string) {
	c := color.New(color.FgHiYellow)
	c.Println(text)
}
