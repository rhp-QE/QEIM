package main

import "time"
import "fmt"

import server "qeim.com/testv1/server"

func main() {
	server.Serve()
	now := time.Now();
	nowInt := uint64(now.Unix())
	fmt.Println(nowInt)


	t := time.Unix(int64(nowInt), 0)
	fmt.Println(t.String())
}