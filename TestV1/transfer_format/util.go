package transfer_format

import (
	"encoding/binary"
	"errors"
	//"fmt"
)

// 提供数据的封包 和 解包工作
// 首部固定4字节表示长度

const (
	modulo     = 65521
	kUint32Len = 4
)

func Packet(message []byte, serviceID uint32) []byte {

	messageLen := len(message)

	//fmt.Println(messageLen)

	result := make([]byte, 3*kUint32Len+messageLen)

	/// 填充总长度
	binary.BigEndian.PutUint32(result, uint32(2*kUint32Len+messageLen))

	/// 填充 serviceID
	binary.BigEndian.PutUint32(result[kUint32Len:], serviceID)

	/// 填充 message
	copy(result[2*kUint32Len:], message)

	/// 计算checkSum 并填充
	checkSum := adler32(result[kUint32Len : 2*kUint32Len+messageLen])
	binary.BigEndian.PutUint32(result[2*kUint32Len+messageLen:], checkSum)

	return result
}

// / message 就是一个包
func UnPacket(message []byte) (*PushMessageObject, error) {
	n := len(message)

	/// 首先检验 checkSun
	_checkSum := adler32(message[:n-kUint32Len])
	checkSum := uint32(binary.BigEndian.Uint32(message[n-kUint32Len:]))
	if checkSum != _checkSum {
		return nil, errors.New("数据包checkSum 有问题")
	}
	//fmt.Println("-----------------------------")
	serviceID := uint32(binary.BigEndian.Uint32(message[:kUint32Len]))
	//fmt.Println("-----------------------------")
	data := message[kUint32Len : n-kUint32Len]
	//fmt.Println("-----------------------------")
	return &PushMessageObject{
		ServiceID: serviceID,
		Data:      data,
		CheckSum:  checkSum,
	}, nil
}

//#pramrk - Private

func adler32(data []byte) uint32 {
	var a, b uint32 = 1, 0

	for _, d := range data {
		a = (a + uint32(d)) % modulo
		b = (b + a) % modulo
	}

	return (b << 16) | a
}
