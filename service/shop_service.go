package service

import (
	"github.com/DaisukeHirabayashi/golang-pagenation/db"
	"github.com/DaisukeHirabayashi/golang-pagenation/dto"
	"github.com/DaisukeHirabayashi/golang-pagenation/entity"
	"github.com/DaisukeHirabayashi/golang-pagenation/mapper"
	"github.com/DaisukeHirabayashi/golang-pagenation/query"
	"github.com/gin-gonic/gin"
)

type ShopService struct{}

var query_pagenation query.Pagenation
var query_sort query.Sort

func (shopService ShopService) GetShops(context *gin.Context) (dto.ShopDto, error) {
	db := db.GetDB()
	var shops []entity.Shop

	totalElements := db.Find(&shops).RowsAffected
	var page dto.Page = mapper.ConvertContextAndTotalElementsToPage(context, int(totalElements))

	if err := db.Scopes(query_pagenation.Pagination(page)).Scopes(query_sort.Sort(context)).Find(&shops).Error; err != nil {
		return dto.ShopDto{}, err
	}

	return dto.ShopDto{Page: page, Shops: shops}, nil
}
