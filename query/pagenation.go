package query

import (
	"github.com/DaisukeHirabayashi/golang-pagenation/dto"
	"github.com/jinzhu/gorm"
)

type Pagenation struct{}

func (pagenation Pagenation) Pagination(page dto.Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page.Number - 1) * page.Size
		return db.Offset(offset).Limit(page.Size)
	}
}
