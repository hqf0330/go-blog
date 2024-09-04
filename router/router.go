package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	// 处理主页面
	http.HandleFunc("/", views.HTML.Index)
	// 数据
	http.HandleFunc("api/v2", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))
}
