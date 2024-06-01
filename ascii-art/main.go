package main

import (
	"fmt"
	"os"

	"ascii/bannerfile"
	"ascii/printascii"
	"ascii/readfile"
	"ascii/splitlines"
	"ascii/escapesequences"
)

func main() {
	// checking if there is enough arguments in the command line
	if len(os.Args) < 2 {
		fmt.Println("Error : Invalid length of arguments")
		os.Exit(0)
	}
	// checks for the banner file in the commad line by calling the GetBannerFile fuction and also 
	// checks if the argument is an empty string
	args := bannerfile.GetBannerFile(os.Args)
	if args == "" {
		return
	}

	// Reading the content of the selected banner file(argss) by calling the ReadFile fuction then
	// storing the variable in the variable initialized(fileContent) while handling the error
	fileContent, err := readfile.ReadFile(args)
	if err != nil {
		fmt.Println("Error: no such named file in the directory")
		return
	}

	// passing the content (fileContent) of the selected banner file (args) through the SplitLines
	// fuction then stores the result into a variable (lines)
	lines := splitlines.SplitLines(fileContent, args)

	arguments := escapesequences.EscapeSequences(os.Args)
	if arguments == "" {
		return
	}

	printascii.PrintASCIIArt(lines, arguments)
}
