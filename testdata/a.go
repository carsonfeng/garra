package testdata

import (
	"github.com/go-xorm/xorm"
)

type Statistic struct {
	RoomId int    `xorm:"'room_id' VARCHAR(32) index" json:"room_id"`
	Dt     string `xorm:"'dt' VARCHAR(32)" json:"dt"`
	Data   string `xorm:"'data' text index" json:"data"`
}

func TestFunc(rids []int) {
	var s *xorm.Session
	rows := make([]*Statistic, 0, len(rids))
	if err := s.In("room_id", rids).Find(&rows); err != nil {
		print("err")
	}
	print("ok")
}
