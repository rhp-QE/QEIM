package roc

import (
	"github.com/go-redis/redis/v8"
)

type SendMessageInterface interface {
	SendMessageToUser(uid uint64, message *Message, complecation func(err error, stataCmd *redis.StatusCmd))
}
