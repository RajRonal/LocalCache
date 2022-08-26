package helper

import (
	"InMemoryCache/database"
	"InMemoryCache/models"
	"github.com/sirupsen/logrus"
)

func InsertData(details models.ToDo) error {
	SQL := `INSERT INTO todo(title,body) VALUES($1,$2)`
	_, err := database.DB.Exec(SQL, details.Title, details.Body)
	if err != nil {
		logrus.Error("InsertData: Error in adding details")
		return err
	}
	return nil
}

func GetData(userId models.UserId) (models.ToDo, error) {
	var result models.ToDo
	SQL := `SELECT title,body from todo where id=$1`
	err := database.DB.Get(&result, SQL, userId.ID)
	if err != nil {
		logrus.Error("GetData: Error in getting details")
		return result, err
	}
	return result, nil
}
