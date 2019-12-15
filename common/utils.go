package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

func (commonError *CommonError) AppendError(key string, err error) {
	commonError.Errors[key] = err.Error()
}
