package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qxsugar/pkg/apix"
	"github.com/qxsugar/pkg/loggerx"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if body, ok := err.(apix.RespBody); ok {
					ctx.PureJSON(body.Code, body)
					ctx.Abort()
				} else if err, ok := err.(error); ok {
					loggerx.GetLogger().Errorw("服务器内部出错", "error", err)
					ctx.PureJSON(apix.Unknown, apix.NewUnknownError(err.(error), ""))
					ctx.Abort()
				} else {
					loggerx.GetLogger().Errorw("服务器内部出错", "error", err)
					ctx.PureJSON(apix.Unknown, apix.NewUnknownError(errors.New(""), fmt.Sprintf("error msg: %v", err)))
					ctx.Abort()
				}
			}
		}()

		ctx.Next()
	}
}
