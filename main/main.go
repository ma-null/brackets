package main

import (
	"bufio"
	"fmt"
	"os"
)

func readStr() (string, error) {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return str, nil
}

var opposite = map[uint8]uint8{
	')': '(',
	'}': '{',
	']': '[',

}

func validateBrackets(str string) bool {
	var stack []byte
	for i := range str {
		switch str[i] {
		case '{', '(', '[':
			stack = append(stack, str[i])
		case '}', ')', ']':
			if len(stack) == 0 || stack[len(stack)-1] != opposite[uint8(str[i])] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	if str, err := readStr(); err == nil {
		fmt.Println(validateBrackets(str))
	} else {
		fmt.Println(err.Error())
	}
}
