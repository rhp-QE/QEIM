package roc

/*
获取用户所连接的服务器地址
*/

type AddrInterface interface {
	UserAddr(uid uint64) string
	AddUserAddr(uid uint64, addr string)
	RemoveUserAddr(uid uint64)

	RedisAddr() string
}
