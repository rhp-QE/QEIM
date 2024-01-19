package main

import (
	"time"

	"github.com/cloudwego/netpoll"
	"github.com/golang/protobuf/proto"
	cofig "qeim.com/testv1/config"
	userpb "qeim.com/testv1/proto/model"
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

		user := &userpb.Person{
			Name: "阮慧鹏",
			Age:  23,
		}

		user2 := &userpb.Person{
			Name: "陈越男",
			Age:  23,
		}

		data, _ := proto.Marshal(user)
		data2, _ := proto.Marshal(user2)

		pc1 := packet.Packet(data, 1)
		pc2 := packet.Packet(data2, 1)

		res := append(pc1, pc2...)

		conn.Write(res)

		//sendMessage(data, conn)
	}
}

func sendMessage(data []byte, connetion netpoll.Connection) {
	msg := packet.Packet(data, 1)
	connetion.Write(msg)
}
