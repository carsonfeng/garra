package a

import (
	"utils"
)

func ban1() {
	_, _ = utils.EmchatSendTxtMsg(0, 0, "", nil)
}

func ban2() string {
	msg, err := utils.EmchatSendTxtMsg(0, 0, "", nil)
	if nil != err {
		return ""
	}
	return msg
}

func ban3() string {
	msg := utils.EmchatSendTxtMsg2(0, 0, "", nil)
	return msg
}

func ban4() string {
	msg := utils.EmchatSendTxtMsg2(0, 0, "", nil)
	msg2 := msg + "ABC"
	return msg2
}

func pass1() string {
	if msg, err := utils.EmchatSendTxtMsg(0, 0, "", nil); nil == err {
		return msg
	}
	return ""
}

func pass2() string {
	go func() {
		msg, err := utils.EmchatSendTxtMsg(0, 0, "", nil)
		print(msg, err)
	}()
	return ""
}
