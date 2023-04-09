package main

import (
	"context"
	"fmt"
	"my-go/grpc/client"
)

func main() {
	c := &client.Client{}
	defer c.Close()
	grpcCli := c.New()
	in := &client.SearchRequest{
		Query:         "nihao",
		PageNumber:    123,
		ResultPerPage: 3,
	}
	res, err := grpcCli.GetList(context.Background(), in)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
