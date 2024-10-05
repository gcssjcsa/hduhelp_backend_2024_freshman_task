package answer

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
	var answers []*models.Answer
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	err = db.GetAnswerListByQid(qid, &answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": answers, "userId": c.Keys["id"].(int)})
}

func Get(c *gin.Context) {
	var answer models.Answer
	aid, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	err = db.GetAnswerById(aid, &answer)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "answer not found"})
		return
	}

	if c.Keys["role"].(models.Role) > models.Role(answer.Permission) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to view this answer"})
		return
	}
	c.JSON(http.StatusOK, answer)
}

func Create(c *gin.Context) {
	var answer models.Answer
	var question models.Question
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if answer.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is empty"})
		return
	}

	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}
	err = db.GetQuestionById(qid, &question)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	answer.QuestionId = qid
	answer.Permission = question.Permission
	postDate := time.Now()
	answer.PostDate = postDate.Format("2006-01-02 15:04:05")
	answer.ModifyDate = answer.PostDate
	answer.Author = c.Keys["username"].(string)
	answer.AuthorId = c.Keys["id"].(int)

	err = db.CreateAnswer(&answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": answer})
}

func Modify(c *gin.Context) {
	var answer, originalAnswer models.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if answer.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is empty"})
		return
	}

	originalAnswer.Id = answer.Id
	err := db.GetAnswerById(answer.Id, &originalAnswer)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "answer not found"})
		return
	} else if c.Keys["id"].(int) != originalAnswer.AuthorId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to modify this answer"})
		return
	}

	now := time.Now()
	answer.ModifyDate = now.Format("2006-01-02 15:04:05")

	err = db.UpdateAnswerById(answer.Id, &answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Update answer successfully"})
}

func Delete(c *gin.Context) {
	var answer models.Answer
	aid, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	err = db.GetAnswerById(aid, &answer)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "answer not found"})
		return
	} else if answer.AuthorId != c.Keys["id"].(int) && c.Keys["role"].(models.Role) != models.Admin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this answer"})
		return
	}

	err = db.DeleteAnswerById(aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Delete answer successfully"})
}
