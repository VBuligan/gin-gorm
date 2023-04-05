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
	var article models.Article
	c.BindJSON(&article)
	err := models.AddNewArticle(&article)
	if err != nil {
		helpers.ResponseJson(c, 404, article)
		return
	}
	helpers.ResponseJson(c, 201, article)
}

func GetArticleById(c *gin.Context) {
	id := c.Param.ByName("id")
	var article models.Article
	err := models.GetArticleById(&article, id)
	if err != nil {
		helpers.ResponseJson(c, 404, article)
		return
	}
	helpers.ResponseJson(c, 200, article)
}

func UpdateArticleById(c *gin.Context) {
	var article models.Article
	id := c.Param.ByName()
	err := models.GetArticleById(&article, id)
	if err != nil {
		helpers.ResponseJson(c, 404, article)
		return
	}
	c.BindJSON(&article)
	err = models.UpdateArticle(&article)
	if err != nil {
		helpers.ResponseJson(c, 404, article)
		return
	}
	helpers.ResponseJson(c, 202, article)
}

func DeleteArticleById(c *gin.Context) {
	var article models.Article
	id := c.Param.ByName("id")
	err := models.DeleteById(&article, id)
	if err != nil {
		helpers.ResponseJson(c, 404, article)
		return
	}
	helpers.ResponseJson(c, 202, article)
}
