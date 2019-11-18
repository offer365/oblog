package main

import (
	"github.com/gin-gonic/gin"
	"github.com/offer365/oblog/controller"
	"github.com/offer365/oblog/dao"
)

func main() {
	router := gin.Default()

	err := dao.Init("117.107.146.194:6603", "oblog", "qwe123!@#QWE", "oblog", "utf8mb4", true)
	if err != nil {
		panic(err)
	}

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("html/*")

	router.GET("/", controller.IndexHandle)
	// 发布文章页面
	router.GET("/article/new/", controller.NewArticle)
	// 文章提交接口
	router.POST("/article/submit/", controller.ArticleSubmit)
	// 文章详情页
	router.GET("/article/detail/", controller.ArticleDetail)

	// 文件上传接口
	router.POST("/upload/file/", controller.UploadFile)

	// 留言页面
	router.GET("/leave/new/", controller.LeaveNew)
	// 关于我页面
	router.GET("/about/me/", controller.AboutMe)

	// 文章评论相关
	router.POST("/comment/submit/", controller.CommentSubmit)

	// 留言相关
	router.POST("/leave/submit/", controller.LeaveSubmit)
	// 分类下面的文章列表
	router.GET("/category/", controller.CategoryList)
	router.Run(":8080")
}
