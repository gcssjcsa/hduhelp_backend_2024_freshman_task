package main

import (
	"MyHelp/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPostQuestionHTML(c *gin.Context) {
	if c.Keys["role"] == models.Guest {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "CUDQuestion.tmpl", "发布")
}

func GetModifyQuestionHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "CUDQuestion.tmpl", "修改")
}
