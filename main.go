package main

import (
	"github.com/VBuligan/gin-gorm/storage"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	storage.DB, err := gorm.Open("postgres", "")
	if err != nil {
		log.Println("Error db accessing", err)
	}
	// * Закрываем соединение
	defer storage.DB.Close()
}
