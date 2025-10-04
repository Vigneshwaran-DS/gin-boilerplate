package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginationResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    200,
		Message: "Success",
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

func SuccessResponseWithPagination(c *gin.Context, data interface{}, page, pageSize, total int) {
	totalPage := (total + pageSize - 1) / pageSize
	c.JSON(200, PaginationResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
		Pagination: Pagination{
			Page:      page,
			PageSize:  pageSize,
			Total:     int64(total),
			TotalPage: totalPage,
		},
	})
}
