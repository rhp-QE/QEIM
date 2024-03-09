package roc

import "sync"

type Roc struct {
	receiveMessageCb func(message *Message)
}

var (
	rocShareInstance *Roc
	onceRoc          sync.Once
)

func RocInstance() *Roc {
	onceRoc.Do(func() {
		rocShareInstance = &Roc{}
	})
	return rocShareInstance
}

// 调用该函数运行
func (roc *Roc) run() {

}

// 设置接受消息回调
func SetReceiveMessageCb()
