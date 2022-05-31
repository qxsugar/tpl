package middleware

import (
	"api-pkg/internal/token"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qxsugar/pkg/apix"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader(token.Header)
		if tokenString != "" {
			uid, err := token.ToUid(tokenString)
			if err == nil {
				ctx.Set(token.UID, uid)
				ctx.Next()
				return
			}
		}

		ctx.JSON(apix.Unauthenticated, apix.NewUnauthenticatedError(errors.New("缺少token"), "缺少token"))
		ctx.Abort()
	}
}
