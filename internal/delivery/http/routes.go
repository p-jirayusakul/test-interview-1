package http

import (
	"github.com/gin-gonic/gin"
)

func BindProductRoutes(group *gin.RouterGroup, productHandler *ProductsHandler) {
	productRouter := group.Group("/product")
	{
		productRouter.POST("/", productHandler.Create)
		productRouter.PATCH("/:id", productHandler.Update)
		productRouter.GET("/:id", productHandler.Get)
	}
}
