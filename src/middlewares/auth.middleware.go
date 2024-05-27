package middlewares

import (
	"ka/src/response"
	"ka/src/services"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a Gin middleware function that performs authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload, err := services.GetUserService().AuthenticateUser(ctx)

		if err != nil {
			response.SendResponse(ctx, payload, payload, err)
			return
		}
		ctx.Next()
	}
}
