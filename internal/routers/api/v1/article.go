package v1

import (
	"gin_blog_service/pkg/app"
	"github.com/gin-gonic/gin"
)

type Article struct{}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context)    {}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
