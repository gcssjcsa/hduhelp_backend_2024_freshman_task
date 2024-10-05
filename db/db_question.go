package db

import (
	"MyHelp/models"
	"log"
)

func GetQuestionById(qid int, q *models.Question) error {
	selectStr := "SELECT * FROM questions WHERE id = ?"
	err := db.QueryRow(selectStr, qid).Scan(&q.Id, &q.Author, &q.Permission, &q.AuthorId, &q.Title, &q.Content,
		&q.PostDate, &q.ModifyDate, &q.Likes)
	if err != nil {
		return err
	}
	return nil
}

func CreateQuestion(q *models.Question) error {
	insertStr := "INSERT INTO questions VALUES (null, ?, ?, ?, ?, ?, ?, ?, 0)"
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

func GetPublicQuestionListByRole(role int, questions *[]*models.Question) error {
	row, err := db.Query("SELECT * FROM questions WHERE permission >= ?", role)
	if err != nil {
		return err
	}

	defer func() {
		if err := row.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}()

	for row.Next() {
		q := new(models.Question)
		err := row.Scan(&q.Id, &q.Author, &q.Permission, &q.AuthorId, &q.Title, &q.Content, &q.PostDate, &q.ModifyDate,
			&q.Likes)
		if err != nil {
			return err
		}
		*questions = append(*questions, q)
	}

	return row.Err()
}

func GetMyQuestionListById(id int, questions *[]*models.Question) error {
	row, err := db.Query("SELECT * FROM questions WHERE id = ?", id)
	if err != nil {
		return err
	}

	defer func() {
		if err := row.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}()

	for row.Next() {
		q := new(models.Question)
		err := row.Scan(&q.Id, &q.Author, &q.Permission, &q.AuthorId, &q.Title, &q.Content, &q.PostDate, &q.ModifyDate,
			&q.Likes)
		if err != nil {
			return err
		}
		*questions = append(*questions, q)
	}

	return row.Err()
}
