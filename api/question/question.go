package question

import (
	"MyHelp/db"
	"MyHelp/models"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetList(c *gin.Context) {

}

func Get(c *gin.Context) {
	var question models.Question
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	err = db.GetQuestionById(qid, &question)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	if int(c.Keys["role"].(models.Role)) > question.Permission {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to view this question"})
		return
	}
	c.JSON(http.StatusOK, question)
}

func Create(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if question.Title == "" || question.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title or Content is empty"})
		return
	}
	postDate := time.Now()
	question.PostDate = postDate.Format("2006-01-02 15:04:05")
	question.ModifyDate = question.PostDate
	question.Author = c.Keys["username"].(string)
	question.AuthorId = int(c.Keys["id"].(float64))

	err := db.CreateQuestion(&question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": question})
}

func Modify(c *gin.Context) {}

func Delete(c *gin.Context) {}
