package server

import (
	"context"
	"encoding/binary"
	"fmt"
	//"go/printer"

	"github.com/cloudwego/netpoll"
	"google.golang.org/protobuf/proto"
	"qeim.com/testv1/config"
	"qeim.com/testv1/log"
	userpb "qeim.com/testv1/pb/generate"
	packet "qeim.com/testv1/transfer_format"
)

//mark const

//#mark var
var logger = log.IMLoggerShareInstance()

var connectionDict map[string]int


func Serve () {
	l, err := netpoll.CreateListener(config.NetConfig.NetWork, config.NetConfig.Address)
	if err != nil {
		logger.Error("创建Listener失败")
		fmt.Println(config.NetConfig.Address)
		fmt.Println(config.NetConfig.NetWork)
		fmt.Println(err)
		return
	}
	connectionDict = make(map[string]int)
	logger.Info("listen over")
	el, _ := netpoll.NewEventLoop(onMessage, netpoll.WithOnPrepare(onNewConnection))
	logger.Info("server start")
	el.Serve(l)
	
}



///#mark Private

/// 当有一个新连接到来时
func onNewConnection(connection netpoll.Connection) (ctx context.Context) {
	logger.Info("new connection")
	logger.Info(connection.RemoteAddr().String())
	connectionDict[connection.RemoteAddr().String()] = 1

	return nil
}


/// 当对方发送新数据时
func onMessage(ctx context.Context, connection netpoll.Connection) error {
	reader := connection.Reader()
	
		pkgLenReader, e1 := reader.Slice(4)
		if e1 != nil {
			println (e1.Error())
		}
		pkgLen := packetLen(pkgLenReader)
		
		pkgReader, e2 := reader.Slice(int(pkgLen))
		if e2 != nil {
			println (e2.Error())
		}

		connection.RemoteAddr()
		uu := connectionDict[connection.RemoteAddr().String()]
		printAddr(connection)
		fmt.Printf("user %d, 发来一条新消息。 如下\n", uu)
		handlePacket(pkgReader)
		return nil
}


func packetLen(reader netpoll.Reader) uint32 {
	defer reader.Release()

	p, _ := reader.ReadBinary(4)
	return uint32(binary.BigEndian.Uint32(p))
}


func handlePacket(reader netpoll.Reader) {
	defer reader.Release()

	data, _ := reader.ReadBinary(reader.Len())
	pushMessage, _ := packet.UnPacket(data)

	if pushMessage.ServiceID == 1 {
		user := &userpb.Person{}

		proto.Unmarshal(pushMessage.Data, user)

		fmt.Print(user.Name)
		fmt.Print("  ")
		fmt.Println(user.Age)
	}
}


func printAddr(connection netpoll.Connection) {
	print(connection.RemoteAddr().String())
}