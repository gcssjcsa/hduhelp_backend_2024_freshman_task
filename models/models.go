package models

import (
	"encoding/json"
	"os"
)

var Conf Config

type Config struct {
	DBUser    string `json:"dbuser"`
	DBPwd     string `json:"dbpwd"`
	DBName    string `json:"dbname"`
	DBHost    string `json:"dbhost"`
	JwtSecret string `json:"jwtsecret"`
}

type Role int

const (
	Admin Role = iota
	Student
	Guest
)

type User struct {
	Id         int
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       Role
	Email      string `json:"email"`
	CreateDate string
}

type Question struct {
	Id         int    `json:"id"`
	Author     string `json:"author"`
	Permission int    `json:"permission"`
	AuthorId   int    `json:"author_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PostDate   string `json:"post_date"`
	ModifyDate string `json:"modify_date"`
	Likes      int    `json:"likes"`
}

type Answer struct {
	Id         int    `json:"id"`
	Author     string `json:"author"`
	Permission int    `json:"permission"`
	AuthorId   int    `json:"author_id"`
	QuestionId int    `json:"question_id"`
	Content    string `json:"content"`
	PostDate   string `json:"post_date"`
	ModifyDate string `json:"modify_date"`
	Likes      int    `json:"like"`
	Dislikes   int    `json:"dislike"`
	IsBest     bool   `json:"is_best"`
}

func init() {
	data, err := os.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	_ = json.Unmarshal(data, &Conf)
}
