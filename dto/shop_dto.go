package dto

import "github.com/DaisukeHirabayashi/golang-pagenation/entity"

type ShopDto struct {
	Page  Page          `json:"page"`
	Shops []entity.Shop `json:"shops"`
}
