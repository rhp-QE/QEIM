package roc

import "sync"

type RocFactory struct {
}

//******************  var  ****************

// 用户地址 组件
var addrComponent AddrInterface

// 发送消息 组件
var sendMessageComponent SendMessageInterface

//*****************************************

var rocFactory *RocFactory
var once sync.Once

func GetRocFactoryShareInstance() *RocFactory {
	once.Do(func() {
		rocFactory = &RocFactory{}
	})
	return rocFactory
}

//*************** GET & SET *****************************

// 用户地址 服务实例
func (factory *RocFactory) GetAddrImpl() AddrInterface {
	if addrComponent == nil {
		addrComponent = NewAddrImpl()
	}
	return addrComponent
}

// 发送消息 组件
func (factory *RocFactory) GetSendMessageImpl() SendMessageInterface {
	if sendMessageComponent == nil {
		sendMessageComponent = NewSendMessageImpl()
	}
	return sendMessageComponent
}

//*******************************************************
