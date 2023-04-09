package main

import (
	"fmt"
	"html/template"
	"log"
	"my-go/study/http/components"
	"my-go/study/http/router"
	"net/http"
	"reflect"
	"strings"
)

func main() {
	components.Init()

	http.HandleFunc("/", handler)
	log.Println("http server start, port: 9090")
	http.ListenAndServe(":9090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var moduleName, controllerName, actionName, routerName = "", "", "", ""
	// 如果是有module的，否则如果没有module的，不符合条件的
	if len(parts) == 3 {
		moduleName, controllerName, actionName = parts[0], parts[1]+"Controller", parts[2]
		routerName = moduleName + "." + controllerName
	} else if len(parts) == 2 {
		_ = moduleName
		routerName = controllerName
	} else {
		notFoundPage(w, r)
		return
	}

	if _, ok := router.WebRouter[routerName]; !ok {
		fmt.Println("不存在的", routerName)
		notFoundPage(w, r)
		return
	}

	fmt.Println(routerName + "." + actionName)
	ctrl := router.WebRouter[routerName]
	controller := reflect.ValueOf(ctrl)
	method := controller.MethodByName(strings.Title(actionName))
	fmt.Println(method)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index"))
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/web/system/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
