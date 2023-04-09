package service

import (
	"context"
	"my-go/grpc/client"
)

func (s *Service) GetList(c context.Context, sr *client.SearchRequest) (list *client.List, err error) {
	list = new(client.List)
	list.Query = sr.Query
	list.PageNumber = sr.PageNumber
	list.ResultPerPage = sr.ResultPerPage
	return
}

func (s *Service) GetLists(sr *client.SearchRequest, sg client.Search_GetListsServer) (err error) {
	return sg.Send(&client.List{
		Query:         sr.Query,
		PageNumber:    sr.PageNumber,
		ResultPerPage: sr.ResultPerPage,
	})
}
