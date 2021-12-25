package main

import (
	"fmt"
	go_say_hello "github.com/igrwijaya/learn-go-module-go-say-hello"
)

func main() {
	hi := go_say_hello.SayHello()

	fmt.Println(hi)
}
