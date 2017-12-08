package service

import (
	"github.com/azenakhi/prog/model"
)

func GetUser() string {
	u := model.UserModel{}
	var m model.Messager
	m = u
	return m.Msg()
}

