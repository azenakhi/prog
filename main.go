package main

import (
	"fmt"
	"prog/service"
)

func main() {
	s := service.GetUser()
	fmt.Println(s)
}