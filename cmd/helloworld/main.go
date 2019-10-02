package main

import (
	"fmt"

	"github.com/devlights/helloworldgo/pkg/helloworld"
)

func main() {
	message, _ := helloworld.Hello(nil)
	fmt.Println(message)
}
