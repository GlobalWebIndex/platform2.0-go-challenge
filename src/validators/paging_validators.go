package validators

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PagingPropsValidator struct {
	PageSize   string `form:"page_size" binding:"required"`
	PageNumber string `form:"page_number" binding:"required"`
}

func NewPagingValidator() PagingPropsValidator {
	return PagingPropsValidator{}
}

func (validator *PagingPropsValidator) GetPagingProps(context *gin.Context) (pageSize uint64, pageNumber uint64, err error) {
	if err = context.Bind(&validator); err != nil {
		err = errors.New("you have to provide page_size and page_number")
		return
	}
	if pageSize, err = strconv.ParseUint(validator.PageSize, 10, 32); err != nil {
		err = errors.New("page_size should be a positive integer")
		return
	}
	if pageNumber, err = strconv.ParseUint(validator.PageNumber, 10, 32); err != nil {
		err = errors.New("page_number should be a positive integer")
		return
	}
	return
}
