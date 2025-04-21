package ui

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	sucessColor = color.New(color.FgGreen).SprintFunc()
	errorColor  = color.New(color.FgRed).SprintFunc()
	warnColor   = color.New(color.FgYellow).SprintFunc()
	infoColor   = color.New(color.FgBlue).SprintFunc()
)

func PrintSucess(msg string) {
	fmt.Println("✅", sucessColor(msg))
}

func PrintError(msg string) {
	fmt.Println("❌", errorColor(msg))
}

func PrintWarning(msg string) {
	fmt.Println("⚠️", warnColor(msg))
}

func PrintInfo(msg string) {
	fmt.Println("ℹ️", infoColor(msg))
}
