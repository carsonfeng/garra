package a

import (
	"dao"
	"models"
)

type name struct {
	A int `json:"a,omitempty"`
}

func TestFuncA() {
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, 3)
	if _, err := dao.Db().Update(&rows); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncB() {
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, 3)
	if _, err := dao.Db().Delete(&rows); err != nil {
		print("err")
	}
	print("ok")
}
