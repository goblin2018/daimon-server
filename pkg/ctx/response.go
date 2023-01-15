package ctx

import (
	"daimon/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func (c *Context) Abort(err error) {
	c.Context.Abort()
	c.JSON(nil, err)
}

type ResponseBody struct {
	Code int32             `json:"code" form:"code"`
	Msg  string            `json:"msg" form:"msg"`
	Data datatypes.JSONMap `json:"data" form:"data"`
}

// TODO 后续优化refreshToken 逻辑
func (c *Context) JSON(data interface{}, err error) {
	refreshToken := c.GetString("refresh_token")

	// 如果请求有错误，数据则为空
	if err != nil {
		data = nil
	}

	er := e.CastError(err)
	c.Context.JSON(http.StatusOK, gin.H{
		"code":          er.Code(),
		"msg":           er.Error(),
		"data":          data,
		"refresh_token": refreshToken,
	})
}

func (c *Context) Fail(err error) {
	c.JSON(nil, err)
}
