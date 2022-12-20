package a

import (
	"dao"
	"github.com/go-xorm/builder"
	"models"
)

type name struct {
	A int `json:"a,omitempty"`
}

func TestFuncA1() {
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, 3)
	rid := 888
	if err := dao.Db().Where("room_id = ?", rid).Find(&rows); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncA2() {
	rids := make([]int, 5, 10)
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, len(rids))
	if err := dao.Db().Where("room_id = ?", 123).In("room_id", rids).Find(&rows); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncA3() {
	rids := make([]int, 5, 10)
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, len(rids))
	uid := 33
	if err := dao.Db().Where(builder.Eq{"room_id": 235, "uid": uid}).In("room_id", []int{1, 2}).Find(&rows); err != nil {
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
func TestFuncWhereAndOr(rids []int) {
	//var s = dao.Db()
	rows := make([]*models.Statistic, 0, len(rids))
	if err := dao.Db().Where("a = ?", 123).And("room_id = ?", 123).Or("room_id = ?", 123).Find(&rows); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncGet(rids []int) {
	var s = dao.Db()
	row := new(models.Statistic)
	if _, err := s.Where("room_id = ?", 123).Get(&row); err != nil {
		print("err")
	}
	print("ok")
}

func TestFuncCount(rids []int) {
	var s = dao.Db()
	row := new(models.Statistic)
	_, _ = s.Where("room_id = ?", 123).Sums(&row)
}
