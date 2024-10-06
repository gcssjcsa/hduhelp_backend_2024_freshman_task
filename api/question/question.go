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

func GetPublicQuesionList(c *gin.Context) {
	var questions []*models.Question
	role := c.Keys["role"].(models.Role)

	err := db.GetPublicQuestionListByRole(int(role), &questions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func GetMyQuesionList(c *gin.Context) {
	var questions []*models.Question
	id := c.Keys["id"].(int)

	err := db.GetMyQuestionListById(id, &questions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func Get(c *gin.Context) {
	var question models.Question
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be integer"})
		return
	}

	err = db.GetQuestionById(qid, &question)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if c.Keys["role"].(models.Role) > models.Role(question.Permission) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to view this question"})
		return
	}

	role := c.Keys["role"].(models.Role)
	if role == models.Guest {
		c.JSON(http.StatusOK, gin.H{"question": question, "userRole": int(role)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"question": question, "userId": c.Keys["id"].(int), "userRole": int(role)})
}

func Create(c *gin.Context) {
	if c.Keys["role"].(models.Role) == models.Guest {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You haven't logged in"})
		return
	}

	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if question.Title == "" || question.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title or content is empty"})
		return
	}
	postDate := time.Now()
	question.PostDate = postDate.Format("2006-01-02 15:04:05")
	question.ModifyDate = question.PostDate
	question.Author = c.Keys["username"].(string)
	question.AuthorId = c.Keys["id"].(int)

	err := db.CreateQuestion(&question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Create question successfully"})
}

func Modify(c *gin.Context) {
	role := c.Keys["role"].(models.Role)
	if role == models.Guest {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You haven't logged in"})
		return
	}

	var question, originalQuestion models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if question.Title == "" || question.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title or content is empty"})
		return
	}

	originalQuestion.Id = question.Id
	err := db.GetQuestionById(question.Id, &originalQuestion)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	} else if c.Keys["id"].(int) != originalQuestion.AuthorId && role != models.Admin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to modify this question"})
		return
	}

	now := time.Now()
	question.ModifyDate = now.Format("2006-01-02 15:04:05")

	err = db.UpdateQuestionById(question.Id, &question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "Update question successfully"})
}

func Delete(c *gin.Context) {
	if c.Keys["role"].(models.Role) == models.Guest {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You haven't logged in"})
		return
	}

	var question models.Question
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be integer"})
		return
	}

	err = db.GetQuestionById(qid, &question)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	} else if question.AuthorId != c.Keys["id"].(int) && c.Keys["role"].(models.Role) != models.Admin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this question"})
		return
	}

	err = db.DeleteQuestionById(qid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Delete question successfully"})
}
