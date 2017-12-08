package service

import (
	"prog/model"
)

func GetUser() string {
	u := model.UserModel{}
	var m model.Messager
	m = u
	return m.Msg()
}

