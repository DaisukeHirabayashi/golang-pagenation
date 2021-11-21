package service

import (
	"github.com/DaisukeHirabayashi/golang-pagenation/db"
	"github.com/DaisukeHirabayashi/golang-pagenation/entity"
	"github.com/gin-gonic/gin"
)

type ShopService struct{}

func (shopService ShopService) GetShops(context *gin.Context) ([]entity.Shop, error) {
	db := db.GetDB()
	var shops []entity.Shop

	if err := db.Find(&shops).Error; err != nil {
		return shops, err
	}
	return shops, nil
}
