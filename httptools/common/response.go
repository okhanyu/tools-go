package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

const (
	CodeError        = -2 // 通用错误
	CodeUnknownError = -1 // 未知错误
	CodeSuccess      = 0
)

// ReplyByError 根据错误返回
func ReplyByError(c *gin.Context, err error) {
	if err != nil {
		FailWithError(c, err)
	} else {
		Success(c, gin.H{})
	}
}

// Success 返回成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    CodeSuccess,
			"message": "success",
			"traceid": "",
			"data":    data,
		},
	)
}

// Fail 返回失败
func Fail(c *gin.Context, code int, message string) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    code,
			"message": message,
			"traceid": "",
		},
	)
}

// FailWithError 根据Error返回失败
func FailWithError(c *gin.Context, err error) {
	// 转换gorm错误
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	FailWithError(c, NewErrorWithCode(RETCOD_NOT_EXISTS))
	//	return
	//}
	//if cErr, ok := err.(BizError); ok {
	//	Fail(c, cErr.Code(), cErr.Error())
	//	return
	//}
	if err != nil {
		Fail(c, CodeError, err.Error())
	} else {
		Fail(c, CodeUnknownError, "unknown")
	}
}
