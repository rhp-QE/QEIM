package model

type DBMessage struct {
	/// 消息类型
	MessageType  int

	/// 消息ID
	MessageID    string
	
	/// 发送者ID
	FormUserID   string

	/// 接受者ID
	ToUserID     string

	/// 消息序号
	MessageSeq   string

	/// 发送时间
	SendTime     string

	/// 消息内容
	Content      string

	/// 所属的群ID
	GroupID      string
}
