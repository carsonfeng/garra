package models

type Statistic struct {
	RoomId int    `xorm:"VARCHAR(32) index" json:"room_id"`
	Uid    int    `xorm:"int index" json:"uid"`
	Dt     string `xorm:"VARCHAR(32)" json:"dt"`
	Data   string `xorm:"text index" json:"data"`
	Empty  string
}
