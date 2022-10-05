package routers

import (
	"github.com/gin-gonic/gin"
	"myGinBlog/utils"
)

func InitRouter() {
	//设置模式，是debug还是其他
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello")
	}

}
