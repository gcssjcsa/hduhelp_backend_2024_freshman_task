package db

import (
	"MyHelp/models"
	"log"
)

func GetAnswerById(aid int, a *models.Answer) error {
	selectStr := "SELECT * FROM answers WHERE id = ?"
	err := db.QueryRow(selectStr, aid).Scan(&a.Id, &a.Author, &a.Permission, &a.AuthorId, &a.QuestionId, &a.Content,
		&a.PostDate, &a.ModifyDate, &a.Likes, &a.Dislikes, &a.IsBest)
	if err != nil {
		return err
	}
	return nil
}

func CreateAnswer(a *models.Answer) error {
	insertStr := "INSERT INTO answers VALUES (null, ?, ?, ?, ?, ?, ?, ?, 0, 0, 0)"
	_, err := db.Exec(insertStr, a.Author, a.Permission, a.AuthorId, a.QuestionId, a.Content, a.PostDate, a.ModifyDate)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAnswerById(aid int, a *models.Answer) error {
	updateStr := "UPDATE answers SET content = ?, permission = ?, modifyDate = ? WHERE id = ?"
	_, err := db.Exec(updateStr, a.Content, a.Permission, a.ModifyDate, aid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAnswerById(aid int) error {
	deleteStr := "DELETE FROM answers WHERE id = ?"
	_, err := db.Exec(deleteStr, aid)
	if err != nil {
		return err
	}
	return nil
}

func GetAnswerListByQid(qid int, answers *[]*models.Answer) error {
	selectStr := "SELECT * FROM answers WHERE question_id = ?"
	rows, err := db.Query(selectStr, qid)
	if err != nil {
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	for rows.Next() {
		a := new(models.Answer)
		err := rows.Scan(&a.Id, &a.Author, &a.Permission, &a.AuthorId, &a.QuestionId, &a.Content,
			&a.PostDate, &a.ModifyDate, &a.Likes, &a.Dislikes, &a.IsBest)
		if err != nil {
			return err
		}
		*answers = append(*answers, a)
	}

	return rows.Err()
}
