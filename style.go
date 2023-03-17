package main

import (
	"github.com/fatih/color"
)

var (
	promptStyle       = color.New(color.FgHiMagenta, color.Bold)
	userInputStyle    = color.New(color.FgHiCyan)
	goGPTResponseStyle = color.New(color.FgHiGreen)
	errorStyle        = color.New(color.FgHiRed, color.Bold)
)

func printStyled(c *color.Color, format string, a ...interface{}) {
	_, _ = c.PrintfFunc()(format, a...)
}

