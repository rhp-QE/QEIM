package transfer_format

//消息进行网络传输的格式
type PushMessageObject struct {
	ServiceID uint32
	Data      []byte
	CheckSum  uint32
}