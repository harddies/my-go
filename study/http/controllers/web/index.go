// Package web web API.
//
// 这是一个swagger 2.0的demo文档
//
// Terms Of Service:
//
// 只是为了demo，不可用来调试哈，这个地址只是自己mvc网站的控制器借个地方来demo swagger的
//
//     Schemes: http, https
//     Host: localhost:9090
//     BasePath: /api/v1
//     Version: 0.0.1
//     License: Ronnie
//     Contact: Ronnie Li<liyixin@bilibili.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package web

import (
	"context"
	"html/template"
	"log"
	"my-go/study/http/components"
	"net/http"
	"time"
)

// swagger:parameters Index Test2
type IndexParam struct {
	// 一个ID
	//
	// Required: true
	// in: path
	Id int `json:"id"`
	// 一个Name
	//
	// Required: true
	// in: query
	Name string `json:"name"`
}

// Index Info
//
// swagger:response IndexResponse
type IndexResponse struct {
	// in: body
	Body ResponseMessage
}

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type IndexController struct {
	controller components.Controller
}

func (ic *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /web/index/index/{id} webIndex Index
	//
	// 通过ID得到一个index response
	//
	// 得到一个ID
	//
	//     Responses:
	//       200: IndexResponse
	t, err := template.ParseFiles("view/web/index/index.html")
	if err != nil {
		log.Fatalln("template error: ", err)
	}
	data := make(map[string]interface{})
	data["name"] = r.FormValue("name")
	data["prices"] = components.Conn.GetAll("select * from py_house_price where id >= ? order by id desc limit ?", 105, 10)
	t.Execute(w, data)
}

func (ic *IndexController) Test(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())

	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		log.Println("goroutine 1 start")
		select {
		case <-ctx.Done():
			log.Println("canceled")
			return
		default:
			log.Println("goroutine 1 done")
			cancel()
		}
	}(ctx)

	go func(ctx context.Context) {
		time.Sleep(3 * time.Second)
		log.Println("goroutine 2 start")
		select {
		case <-ctx.Done():
			log.Println("canceled")
			return
		default:
			log.Println("goroutine 2 done")
		}
	}(ctx)

	time.Sleep(2 * time.Second)
}
