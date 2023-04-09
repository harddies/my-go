package router

import (
	"my-go/study/http/controllers/web"
)

var WebRouter map[string]interface{}

func init() {
	WebRouter = make(map[string]interface{})
	WebRouter["web.indexController"] = &web.IndexController{}
}
