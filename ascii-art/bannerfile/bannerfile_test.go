package bannerfile

import (
	"fmt"
	"testing"
)

func TestGetBannerFile(t *testing.T) {
	args := []string{"Program", "arg2", "standard"}
	expected := "standard.txt"
	fileName := GetBannerFile(args)
	if fileName != expected {
		fmt.Println("not expected argument")
	}
}
