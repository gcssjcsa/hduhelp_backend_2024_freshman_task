package db

import "MyHelp/models"

func GetQuestionById(qid int, q *models.Question) error {
	selectStr := "SELECT * FROM questions WHERE id = ?"
	err := db.QueryRow(selectStr, qid).Scan(&q.Id, &q.Author, &q.Permission, &q.AuthorId, &q.Title, &q.Content,
		&q.PostDate, &q.ModifyDate, &q.Likes, &q.IsBest)
	if err != nil {
		return err
	}
	return nil
}

func CreateQuestion(q *models.Question) error {
	insertStr := "INSERT INTO questions VALUES (null, ?, ?, ?, ?, ?, ?, ?, 0, 0)"
	_, err := db.Exec(insertStr, q.Author, q.Permission, q.AuthorId, q.Title, q.Content, q.PostDate, q.ModifyDate)
	if err != nil {
		return err
	}
	return nil
}

func UpdateQuestionById(qid int, q *models.Question) error {
	updateStr := "UPDATE questions SET title = ?, content = ?, permission = ?, modifyDate = ? WHERE id = ?"
	_, err := db.Exec(updateStr, q.Title, q.Content, q.Permission, q.ModifyDate, qid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteQuestionById(qid int) error {
	deleteStr := "DELETE FROM questions WHERE id = ?"
	_, err := db.Exec(deleteStr, qid)
	if err != nil {
		return err
	}
	return nil
}

func GetPublicQuestionList() ([]models.Question, error) {
	
}