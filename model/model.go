package model

import (
	"fmt"
)

type UserModel struct {
	name string
	age string
}

type Messager interface {
	Msg() string
}

func (u UserModel) Msg() string {
	defer fmt.Println("hello")
	u1 := UserModel{name:"Ahmed",age:"40"}
	return u1.age
}