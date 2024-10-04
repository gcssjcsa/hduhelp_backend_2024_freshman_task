package main

import (
	"MyHelp/api/answer"
	"MyHelp/api/comment"
	"MyHelp/api/question"
	"MyHelp/api/user"
	"MyHelp/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/src", "./static")
	r.LoadHTMLGlob("templates/*")

	frontend := r.Group("/")
	{
		frontend.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		frontend.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", nil)
		})
		frontend.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})
		frontend.GET("/user", func(c *gin.Context) {
			c.HTML(http.StatusOK, "user.html", nil)
		})
		frontend.GET("/question/:qid", func(c *gin.Context) {
			c.HTML(http.StatusOK, "question.html", nil)
		})
		frontend.GET("/question/post", user.AuthMiddleware(), GetPostQuestionHTML)
		frontend.GET("/question/:qid/modify", user.AuthMiddleware(), GetModifyQuestionHTML)
	}

	// 注册账户
	r.POST("/register", user.Register)
	// 登录账户
	r.POST("/login", user.Login)
	// 退出账户
	r.DELETE("/login", user.Logout)
	// 更新、删除用户信息
	r.POST("/user", user.AuthMiddleware(), user.GetProfile)
	r.PUT("/user", user.AuthMiddleware(), user.UpdateProfile)
	r.PATCH("/user", user.AuthMiddleware(), user.UpdatePassword)
	r.DELETE("/user", user.AuthMiddleware(), user.Delete)

	// 内容API
	apiRoute := r.Group("/api")
	{
		apiRoute.Use(user.AuthMiddleware())

		// 问题内容
		apiRoute.GET("/question", question.GetList)
		apiRoute.POST("/question", question.Create)
		apiRoute.GET("/question/:qid", question.Get)
		apiRoute.PUT("/question/:qid", question.Modify)
		apiRoute.DELETE("/question/:qid", question.Delete)

		// 问题回答
		apiRoute.GET("/question/:qid/answer", answer.GetList)
		apiRoute.POST("/question/:qid/answer/", answer.Create)
		apiRoute.GET("/question/:qid/answer/:aid", answer.Get)
		apiRoute.PUT("/question/:qid/answer/:aid", answer.Modify)
		apiRoute.DELETE("/question/:qid/answer/:aid", answer.Delete)

		// 回答评论
		apiRoute.GET("/comment", comment.Get)
		apiRoute.POST("/comment", comment.Send)
		apiRoute.PUT("/comment", comment.Modify)
		apiRoute.DELETE("/comment", comment.Delete)
	}

	err := r.Run()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()
}
