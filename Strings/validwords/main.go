package main

import (
	"fmt"
	"valid/validate"
)

func main() {
	result := validate.IsValid("Hello92")
	fmt.Println(result)
}
