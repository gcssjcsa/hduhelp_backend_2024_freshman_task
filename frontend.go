package main

import (
	"MyHelp/db"
	"MyHelp/models"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPostQuestionHTML(c *gin.Context) {
	if c.Keys["role"] == models.Guest {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "postQuestion.tmpl", "发布")
}

func GetModifyQuestionHTML(c *gin.Context) {
	var q models.Question
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	err = db.GetQuestionById(qid, &q)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	} else if c.Keys["id"].(int) != q.AuthorId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not author"})
		return
	}

	c.HTML(http.StatusOK, "postQuestion.tmpl", "修改")
}
