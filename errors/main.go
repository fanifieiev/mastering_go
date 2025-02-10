package main

import (
	"errors"
	"fmt"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("Both values are 0")
	}
	return nil
}

func check2(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("Both values are 0")
	}
	return nil
}

func main() {
	err := check(0, 10)
	if err != nil {
		fmt.Println(err)
	}
}
