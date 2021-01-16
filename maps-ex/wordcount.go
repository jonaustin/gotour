package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	strmap := make(map[string]int)

	for i := 0; i < len(strs); i++ {
		strmap[strs[i]] = len(strs[i])
	}

	return strmap
}

func main() {
	fmt.Println(WordCount("fook bar ba"))
	wc.Test(WordCount)
}
