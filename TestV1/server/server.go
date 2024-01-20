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
	pbMessage "qeim.com/testv1/pb/generate"
	packet "qeim.com/testv1/transfer_format"
	pushMessageHandle "qeim.com/testv1/server/handler"

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
	return nil
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
		imCloudMessageResp := pushMessageHandle.Handle(imCloudMessage, con)

		dataResp, _ := proto.Marshal(imCloudMessageResp)
		packet := packet.Packet(dataResp, 1)
		con.Writer().WriteBinary(packet)
	}
}


func printAddr(connection netpoll.Connection) {
	print(connection.RemoteAddr().String())
}