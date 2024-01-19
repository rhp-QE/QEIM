package main

import (
	"context"
	"fmt"
	netpoll "github.com/cloudwego/netpoll"
	server "qeim.com/testv1/server"
)

func main() {
	server.Serve()
}


///#pramrk private
func onRequest(ctx context.Context, connection netpoll.Connection) error {
	arr := make([]byte, 1024)
	connection.Read(arr)
	
	fmt.Println("read from connect : arr")
	fmt.Println(arr)
	return nil
}

// NOTE: Implement
func prepare(connection netpoll.Connection) (ctx context.Context) {
	fmt.Println("12312")
	return nil;
}
