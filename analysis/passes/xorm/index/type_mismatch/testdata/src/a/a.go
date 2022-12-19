package a

import (
	"dao"
	"models"
)

//func TestFunc2(rids [3]string) {
//	var s *xorm.Session
//	rows := make([]*models.Statistic, 0, len(rids))
//	if err := s.In("room_id", rids).Find(&rows); err != nil {
//		print("err")
//	}
//	print("ok")
//}

type name struct {
	A int `json:"a,omitempty"`
}

func TestFunc() {
	rids := make([]int, 5, 10)
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, len(rids))
	if err := dao.Db().Where("room_id = ?", 123).In("room_id", rids).Find(&rows); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncB(rids []int) {
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, len(rids))
	if err := dao.Db().Where("room_id = ?", 123).NotIn("room_id", rids).Find(&rows); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncC(uids []string) {
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, len(uids))
	if err := dao.Db().In("uid", uids).Find(&rows); err != nil {
		print("err")
	}
	print("ok")
}
