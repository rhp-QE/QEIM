package roc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type sendMessageImpl struct {
}

// **** New *****
func NewSendMessageImpl() *sendMessageImpl {
	return &sendMessageImpl{}
}

//**************

// **** Impl sendMessageInterface *****
func (sender *sendMessageImpl) SendMessageToUser(uid uint64, message *Message, complecation func(err error, stataCmd *redis.StatusCmd)) error {
	addrInterface := GetRocFactoryShareInstance().GetAddrImpl()

	rdb := redis.NewClient(&redis.Options{
		Addr:     addrInterface.RedisAddr(),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	if messageData, err := TransferMessageToBytes(message); err != nil {
		go func() {
			res := rdb.Set(ctx, UserMessageBoxKey(uid), messageData, 0)
			complecation(nil, res)
		}()
		rdb.Set(ctx, UserMessageBoxKey(uid), messageData, 0)
	} else {
		fmt.Printf("[error]:%s", err.Error())
	}

	return nil
}

//*************************************
