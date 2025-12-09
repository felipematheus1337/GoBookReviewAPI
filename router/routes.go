package router

import (
	"github.com/felipematheus1337/GoBookReviewAPI/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/override/docs"
)

func InitializeRoutes(router *gin.Engine, b *handler.BookHandler, r *handler.ReviewHandler) {

	basePath := "api/v1/book"

	docs.SwaggerInfo.BasePath = basePath

	v1Books := router.Group(basePath)

	v1Reviews := router.Group(basePath)

	RegisterBookRoutes(v1Books, b)

	RegisterReviewRoutes(v1Reviews, r)

}

func RegisterBookRoutes(v1 *gin.RouterGroup, b *handler.BookHandler) {
	{
		v1.POST("/", b.Create)
		v1.GET("/", b.List)
	}
}

func RegisterReviewRoutes(v1 *gin.RouterGroup, r *handler.ReviewHandler) {
	{
		v1.POST("/", r.Create)
		v1.GET("/", r.List)
	}
}
