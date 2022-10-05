package routers

import (
	"github.com/gin-gonic/gin"
	v1 "myGinBlog/api/v1"
	"myGinBlog/utils"
)

func InitRouter() {
	//设置模式，是debug还是其他
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	///*
	//	后台管理路由接口
	//*/
	//auth := r.Group("api/v1")
	//auth.Use(middleware.JwtToken())

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		// 文章分类信息模块
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)
		router.POST("category/add", v1.AddCategory)
		router.PUT("category/:id", v1.EditCate)
		router.DELETE("category/:id", v1.DeleteCate)

		// 文章模块
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.GET("article/:id", v1.EditArt)
		router.POST("article/add", v1.AddArticle)
		router.DELETE("article/:id", v1.DeleteArt)

		//
		//	// 登录控制模块
		//	router.POST("login", v1.Login)
		//	router.POST("loginfront", v1.LoginFront)
		//
		//	// 获取个人设置信息
		//	router.GET("profile/:id", v1.GetProfile)
		//
		//	// 评论模块
		//	router.POST("addcomment", v1.AddComment)
		//	router.GET("comment/info/:id", v1.GetComment)
		//	router.GET("commentfront/:id", v1.GetCommentListFront)
		//	router.GET("commentcount/:id", v1.GetCommentCount)
	}

	_ = r.Run(utils.HttpPort)

}
