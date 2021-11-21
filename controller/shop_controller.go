package controller

import (
	"net/http"

	"github.com/DaisukeHirabayashi/golang-pagenation/service"
	"github.com/gin-gonic/gin"
)

type ShopController struct{}

var shopService service.ShopService

func (shopController ShopController) GetShops(context *gin.Context) {
	shops, err := shopService.GetShops(context)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
	} else {
		context.JSON(http.StatusOK, shops)
	}
}
