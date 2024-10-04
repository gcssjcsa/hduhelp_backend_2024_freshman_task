package db

import "MyHelp/models"

func InsertNewUserRecord(user models.User) error {
	insertStr := "INSERT INTO users (username, password, email, createDate) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(insertStr, user.Username, user.Password, user.Email, user.CreateDate)
	if err != nil {
		return err
	}
	return nil
}

func SelectLoginUserPassword(loginUser *models.User, user *models.User) error {
	selectStr := "SELECT password FROM users WHERE username = ?"
	return db.QueryRow(selectStr, loginUser.Username).Scan(&user.Password)
}

func GetLoginUserInfoByName(username string, save *models.User) error {
	selectStr := "SELECT id, username, role FROM users WHERE username = ?"
	return db.QueryRow(selectStr, username).Scan(&save.Id, &save.Username, &save.Role)
}

func SelectUserProfileById(user *models.User) error {
	selectStr := "SELECT username, role, email FROM users WHERE id = ?"
	return db.QueryRow(selectStr, user.Id).Scan(&user.Username, &user.Role, &user.Email)
}

func UpdateUserProfile(newUserInfo *models.User) error {
	updateStr := "UPDATE users SET username = ?, email = ? WHERE id = ?"
	_, err := db.Exec(updateStr, newUserInfo.Username, newUserInfo.Email, newUserInfo.Id)
	if err != nil {
		return err
	}
	return nil
}

func SelectUserPassword(userid int) (string, error) {
	var password string
	selectStr := "SELECT password FROM users WHERE id = ?"
	err := db.QueryRow(selectStr, userid).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func UpdateUserPassword(userid int, newPwd string) error {
	updateStr := "UPDATE users SET password = ? WHERE id = ?"
	_, err := db.Exec(updateStr, newPwd, userid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userid int) error {
	deleteStr := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(deleteStr, userid)
	if err != nil {
		return err
	}
	return nil
}
