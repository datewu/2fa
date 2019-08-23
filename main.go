package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	tag := randStr()
	fmt.Println("The QR CODE:", strings.ToUpper(tag))
	for {
		code, remain := Gen2fa(tag, 30) // default interval is 30 second
		for remain > 0 {
			fmt.Printf("code: %06d (%d second(s))\r", code, remain)
			time.Sleep(1 * time.Second)
			remain--
		}
	}
}
