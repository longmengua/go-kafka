package tools

import (
	"fmt"
)

func Async(goFunc func()) {
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		goFunc()
	}()
}
