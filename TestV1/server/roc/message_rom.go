package roc

import "encoding/json"

type Message struct {
	FromUid        uint64
	ToUid          uint64
	data           []byte
	sendTime       string
	conversationID string
}

func TransferMessageToBytes(data *Message) ([]byte, error) {
	var jsonData []byte
	var err error

	if jsonData, err = json.Marshal(data); err != nil {
		return jsonData, err
	}

	return jsonData, nil
}

func TransferBytesToMessage(data []byte) (*Message, error) {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
