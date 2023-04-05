package routers

import (
	"github.com/VBuligan/gin-gorm/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // * Аналог mux.NewRouter

	// * В gin принято группировать ресурсы
	apiV1Group := router.Group("/api/v1") // * prefix
	{
		apiV1Group.GET("article", handlers.GetAllArticles)
		apiV1Group.POST("article", handlers.PostNewArticles)
		apiV1Group.GET("article/:id", handlers.GetArticleById)
		apiV1Group.PUT("article/:id", handlers.UpdateArticleById)
		apiV1Group.DELETE("article/:id", handlers.DeleteArticleById)
	}

	return router
}
