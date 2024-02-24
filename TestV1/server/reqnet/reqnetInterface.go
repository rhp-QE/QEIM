package reqnet

type reqnetInterface interface {
	// 发送消息 （目前是发送到消息队列)
	SendMessage(uid uint64, msg []byte)

	// 注入接受到消息的函数
	InjectReceiveMessageHandle(func(uid uint64, msg []byte))

	// 开始
	Start()

	// 终止
	Stop()
}
