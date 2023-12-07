package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "치환할문자열"
	fmt.Println(strings.Replace(str, "치환할", "치환된", -1))
}
