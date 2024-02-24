package local

import "github.com/cloudwego/netpoll"

var (
	/// 用户uid 到connection的映射
	uidToConnection map[uint64]netpoll.Connection = make(map[uint64]netpoll.Connection)
)

func ConecntionForUid(uid uint64) netpoll.Connection {
	return uidToConnection[uid]
}

func StoreConnection(uid uint64, con netpoll.Connection) {
	uidToConnection[uid] = con
}
