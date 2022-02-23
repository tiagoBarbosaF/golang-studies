package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown foz jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v0\n", doubleRev, doubleRevErr)
}

func Reverse(var1 string) (string, error) {
	if !utf8.ValidString(var1) {
		return var1, errors.New("input is not valid UTF-8")
	}
	var2 := []rune(var1)

	for var3, var4 := 0, len(var2)-1; var3 < len(var2)/2; var3, var4 = var3+1, var4-1 {
		var2[var3], var2[var4] = var2[var4], var2[var3]
	}
	return string(var2), nil
}
