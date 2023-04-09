package client

import (
	"google.golang.org/grpc"
	"log"
)

type Client struct {
}

var conn *grpc.ClientConn

func (c *Client) New() SearchClient {
	var err error
	conn, err = grpc.Dial("localhost:7999", grpc.WithInsecure())
	if err != nil {
		panic("dial failed")
	}
	log.Println("grpc连接成功")
	client := NewSearchClient(conn)
	log.Println("SearchClient创建成功")
	return client
}

func (c *Client) Close() {
	conn.Close()
}
