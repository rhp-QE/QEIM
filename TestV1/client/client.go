package main

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/cloudwego/netpoll"
	"github.com/golang/protobuf/proto"
	cofig "qeim.com/testv1/config"
	pbMessage "qeim.com/testv1/pb/generate"
	packet "qeim.com/testv1/transfer_format"
)

func main() {
	// Dial a connection with Dialer.
	dialer := netpoll.NewDialer()
	conn, err := dialer.DialConnection(cofig.NetConfig.NetWork, cofig.NetConfig.Address, time.Second)
	if err != nil {
		panic("dial netpoll connection failed")
	}

	defer conn.Close()
	for {
		time.Sleep(time.Second)

		msg := &pbMessage.IMCloudPbMessage{
			CmdType: pbMessage.Cmd_RegisterCmd,
			IsRequest: true,
			RequestBody: &pbMessage.IMCloudPbMessageRequestBody{
				RegisterRequestBody: &pbMessage.RegisterRequest{
					UserName: "RuanHuipeng",
					Password: "12345678",
				},
			},
		}

		data, _ := proto.Marshal(msg)

		pc1 := packet.Packet(data, 1)

		conn.Write(pc1)

	}
}

func sendMessage(data []byte, connetion netpoll.Connection) {
	msg := packet.Packet(data, 1)
	connetion.Write(msg)
}


func handleRead(connection netpoll.Connection) {
	reader := connection.Reader()
	
	for {
		pkgLenReader, error1 := reader.Slice(4)
		if error1 != nil {
			println (error1.Error())
		}
		pkgLen := packetLen(pkgLenReader)
		
		pkgReader, error2 := reader.Slice(int(pkgLen))
		if error2 != nil {
			println (error2.Error())
		}

		//uu := connectionDict[connection.RemoteAddr().String()]
		//printAddr(connection)
		fmt.Printf("user %d, 发来一条新消息。\n", connection)
		handlePacket(connection, pkgReader)
	}
}

func printMessage(pushMessage *pbMessage.IMCloudPbMessage) {
	if pushMessage == nil {
		return
	}

	if pushMessage.CmdType == pbMessage.Cmd_LoginCmd {
		fmt.Println(pushMessage.ResponseBody.LoginResponseBody.StringState)
	}

	if pushMessage.CmdType == pbMessage.Cmd_RegisterCmd {
		fmt.Println(pushMessage.ResponseBody.RegisterResponseBody.StringState)
	}
}


func packetLen(reader netpoll.Reader) uint32 {
	defer reader.Release()

	p, _ := reader.ReadBinary(4)
	return uint32(binary.BigEndian.Uint32(p))
}


func handlePacket(con netpoll.Connection, reader netpoll.Reader) {
	defer reader.Release()

	data, _ := reader.ReadBinary(reader.Len())
	pushMessage, _ := packet.UnPacket(data)

	if pushMessage.ServiceID == 1 {
		imCloudMessage := &pbMessage.IMCloudPbMessage{}

		proto.Unmarshal(pushMessage.Data, imCloudMessage)
		printMessage(imCloudMessage)
	}
}