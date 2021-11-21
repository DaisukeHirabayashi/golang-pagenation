package server

import (
	"github.com/DaisukeHirabayashi/golang-pagenation/controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	engine := router()

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

func router() *gin.Engine {
	router := gin.Default()

	// Routing
	shopRouter := router.Group("/shops")
	{
		shopController := controller.ShopController{}
		shopRouter.GET("", shopController.GetShops)
	}

	return router
}
