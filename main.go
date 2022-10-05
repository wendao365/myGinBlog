package main

import (
	"myGinBlog/model"
	"myGinBlog/routers"
)

func main() {
	// 引用数据库
	model.InitDB()
	// 引入路由组件
	routers.InitRouter()
}
