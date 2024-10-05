package user

import (
	"MyHelp/db"
	"MyHelp/models"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"time"
	"unicode/utf8"
)

func Register(c *gin.Context) {
	var user, existedUser models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
		return
	}

	if utf8.RuneCountInString(user.Username) > 63 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is too long"})
	}
	if utf8.RuneCountInString(user.Email) > 63 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is too long"})
	}

	matched, err := regexp.Match("^[A-Za-z0-9]+@[A-Za-z0-9]+\\.[A-Za-z0-9]+$", []byte(user.Email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is invalid"})
		return
	}

	err = db.GetLoginUserInfoByName(user.Username, &existedUser)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	} else if !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	nowTime := time.Now()
	user.CreateDate = nowTime.Format("2006-01-02 15:04:05")

	err = db.InsertNewUserRecord(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Register sucessfully!"})
}

func Login(c *gin.Context) {
	var user, loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginUser.Username == "" || loginUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
		return
	}

	err := db.SelectLoginUserPassword(&loginUser, &user)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		// 没有此用户
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		// 密码错误
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	err = db.GetLoginUserInfoByName(loginUser.Username, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenString, err := token.SignedString([]byte(models.Conf.JwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't generate token"})
		return
	}

	c.SetCookie("token", tokenString, 3600*24*7, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"Status": "Login sucessfully!", "token": tokenString})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"Status": "Logout sucessfully!"})
}

func GetProfile(c *gin.Context) {
	var user models.User
	role := c.Keys["role"].(models.Role)
	if role == models.Guest {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You haven't logged in", "role": int(role)})
		return
	} else {
		user.Id = c.Keys["id"].(int)
		err := db.SelectUserProfileById(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": user.Id, "username": user.Username, "role": user.Role, "email": user.Email})
	}
}

func UpdateProfile(c *gin.Context) {
	var newUserInfo, existedUser models.User
	newUserInfo.Id = c.Keys["id"].(int)

	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUserInfo.Username == "" || newUserInfo.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Info can not be empty"})
		return
	}

	if utf8.RuneCountInString(newUserInfo.Username) > 63 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is too long"})
	}
	if utf8.RuneCountInString(newUserInfo.Email) > 63 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is too long"})
	}

	matched, err := regexp.Match("^[A-Za-z0-9]+@[A-Za-z0-9]+\\.[A-Za-z0-9]+$", []byte(newUserInfo.Email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is invalid"})
		return
	}

	err = db.GetLoginUserInfoByName(newUserInfo.Username, &existedUser)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	} else if !errors.Is(err, sql.ErrNoRows) && existedUser.Id != newUserInfo.Id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	err = db.UpdateUserProfile(&newUserInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Update sucessfully!"})
}

func UpdatePassword(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var Data map[string]interface{}
	err = json.Unmarshal(rawData, &Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	originalPwd := Data["originalPwd"].(string)
	newPwd := Data["newPwd"].(string)
	id := c.Keys["id"].(int)

	if password, err := db.SelectUserPassword(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	} else if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(originalPwd)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	} else {
		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(newPwd), bcrypt.DefaultCost)
		newPwd = string(hashedPwd)
		err = db.UpdateUserPassword(id, newPwd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// 清除登录状态
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"Status": "Change password sucessfully!"})
	}
}

func Delete(c *gin.Context) {
	id := c.Keys["id"].(int)
	err := db.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"Status": "Delete sucessfully! Bye!"}) // 重定向由前端完成
}
