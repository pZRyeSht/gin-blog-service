package routers

import (
	_ "gin_blog_service/docs"
	"gin_blog_service/internal/middleware"
	"gin_blog_service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	article := v1.NewArticle()
	tag := v1.NewTag()

	api := r.Group("/api/v1")
	{
		api.POST("/tags", tag.Create)
		api.DELETE("/tags/:id", tag.Delete)
		api.PUT("/tags/:id", tag.Update)
		api.PATCH("/tags/:id/state", tag.Update)
		api.GET("/tags", tag.List)

		api.POST("/articles", article.Create)
		api.DELETE("/articles/:id", article.Delete)
		api.PUT("/articles/:id", article.Update)
		api.PATCH("/articles/:id/state", article.Update)
		api.GET("/articles/:id", article.Get)
		api.GET("/articles", article.List)
	}

	return r
}
