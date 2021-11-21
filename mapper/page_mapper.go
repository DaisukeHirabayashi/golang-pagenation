package mapper

import (
	"log"
	"math"
	"strconv"

	"github.com/DaisukeHirabayashi/golang-pagenation/dto"
	"github.com/gin-gonic/gin"
)

func ConvertContextAndTotalElementsToPage(context *gin.Context, totalElements int) dto.Page {
	page, _ := strconv.Atoi(context.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(context.Query("size"))
	log.Print(pageSize)
	log.Print(totalElements)
	switch {
	case pageSize > totalElements:
		pageSize = totalElements
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		if totalElements < 5 {
			pageSize = totalElements
		} else {
			pageSize = 5
		}
	}
	totalPages := int(math.Ceil(float64(totalElements) / float64(pageSize)))

	return dto.Page{Number: page, Size: pageSize, TotalElements: totalElements, TotalPages: totalPages}
}
