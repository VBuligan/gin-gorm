package handlers

import (
	"github.com/VBuligan/gin-gorm/helpers"
	"github.com/VBuligan/gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	err := models.GetAllArticles(&articles)
	if err != nil {
		helpers.ResponseJson(c, 404, articles)
		return
	}
	helpers.ResponseJson(c, 200, articles)
}

func PostNewArticles(c *gin.Context) {

}
func GetArticleById(c *gin.Context) {

}
func UpdateArticleById(c *gin.Context) {

}
func DeleteArticleById(c *gin.Context) {

}
