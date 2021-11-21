package query

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Sort struct{}

func (sort Sort) Sort(context *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		direction := context.Query("direction")
		if direction != "asc" && direction != "desc" {
			direction = "desc"
		}

		orderby := context.Query("orderby")
		if orderby == "" {
			orderby = "id"
		}

		order := strings.Join([]string{orderby, direction}, " ")
		return db.Order(order)
	}
}
