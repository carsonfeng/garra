package a

import (
	"dao"
	"models"
)

type Temp struct {
}

func (t *Temp) Update(b interface{}) {
	return
}

func TestFuncRejectA() {
	//var s = dao.Db()
	if _, err := dao.Db().Update(&models.Statistic{}); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncRejectB() {
	//var s = dao.Db()
	if _, err := dao.Db().Delete(&models.Statistic{}); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncRejectC() {
	session := dao.Db()
	session.Update(&models.Statistic{})
}

func TestFuncPassC() {
	session := dao.Db().Id(1)
	session.Update(&models.Statistic{})
}

func TestFuncPassD() {
	t := Temp{}
	t.Update(&models.Statistic{})
}

func TestFuncPassE() {
	session := dao.Db().Id(1)
	session = session.Incr("fail", 1)
	session.Update(&models.Statistic{})
}
