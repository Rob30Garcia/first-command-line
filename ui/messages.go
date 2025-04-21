package ui

import (
	"github.com/fatih/color"
)

var (
	sucessColor = color.New(color.FgGreen).SprintFunc()
	errorColor  = color.New(color.FgRed).SprintFunc()
	warnColor   = color.New(color.FgYellow).SprintFunc()
	infoColor   = color.New(color.FgBlue).SprintFunc()
)

func PrintSucess(msg string) string {
	return sucessColor(msg)
}

func PrintError(msg string) string {
	return errorColor(msg)
}

func PrintWarning(msg string) string {
	return warnColor(msg)
}

func PrintInfo(msg string) string {
	return infoColor(msg)
}
