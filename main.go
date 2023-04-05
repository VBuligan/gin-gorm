package main

import (
	"github.com/VBuligan/gin-gorm/models"
	"github.com/VBuligan/gin-gorm/routers"
	"github.com/VBuligan/gin-gorm/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	var err error
	storage.DB, err = gorm.Open("postgres", "host=... user=... password=... dbname=...")
	if err != nil {
		log.Println("Error db accessing", err)
	}
	// * Закрываем соединение
	defer storage.DB.Close()

	// * В этот момент орм сама сгенерит все запросы, миграции и применит их
	storage.DB.AutoMigrate(&models.Article{})

	// * r - gin маршрутизатор
	r := routers.SetupRouter()
	r.Run()
}
