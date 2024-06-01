package bannerfile

import (
	"fmt"
)

// get the filename from the command line arguments
func GetBannerFile(args []string) string {
	if len(args) == 3 {
		if args[2] == "shadow" {
			return "shadow.txt"
		}
		if args[2] == "thinkertoy" || args[2] == "thinkertoy.txt" {
			return "thinkertoy.txt"
		}
		if args[2] == "standard" || args[2] == "standard.txt" {
			return "standard.txt"
		} else {
			fmt.Println(`Unidentified file, Usage : go run . "desiredtext" standard.txt`)
			return ""
		}
	} else {
		return "standard.txt"
	}
}
