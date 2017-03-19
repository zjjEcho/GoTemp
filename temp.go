package main

import (
	// "bytes"
	"fmt"
	"strings"
)

func laststep(r rune) bool {
	switch r {
	case ':':
		return true
	case '_':
		return true
	case '-':
		return true
	case ' ':
		return true
	}

	return false
}

func main() {
	// fmt.Println(strings.ContainsAny("failure", "fb"))
	// fmt.Println(strings.SplitN("a:b:c:d", ":", -1))
	// fmt.Println(strings.SplitAfterN("a:b:c:d", ":", -1))
	// fmt.Println(strings.LastIndexAny("a:b_c-d", "_"))
	// keyStr := "a:ni-xx 不是吧_test:哈哈"
	// keyStr := "rider:shop:12344"
	keyStr := "RC_SHOP_36_test_33455"
	index := strings.LastIndexFunc(keyStr, laststep)
	s := []byte(keyStr)
	str := string(s[:index])
	fmt.Println(str)
}
