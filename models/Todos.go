package models

import (
	"example.com/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}
