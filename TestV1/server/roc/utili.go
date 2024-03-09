package roc

import "strconv"

// 工具方法

const (
	messageBoxPre = "messageBox?"
)

func UserMessageBoxKey(uid uint64) string {
	str1 := strconv.FormatUint(uid, 10)
	return messageBoxPre + str1
}
