package main

import (
	"fmt"
	"github.com/azenakhi/prog/service"
)

func main() {
	s := service.GetUser()
	fmt.Println(s)
}
