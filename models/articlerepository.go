package models

import "github.com/VBuligan/gin-gorm/storage"

func GetAllArticles(a *[]Article) error {
	return storage.DB.Find(a).Error
}

func AddNewArticle(a *Article) error {
	return storage.DB.Create(a).Error
}

func AddNewArticle() {

}

func AddNewArticle() {

}

func AddNewArticle() {

}
