package middlewares

import (
	"github.com/gin-gonic/gin"
	"gitlab.devgroup.tech/shkolkovo/romanych/api/v1/response"
)

func UserMiddleware(c *gin.Context) {
	var err error
	var req struct {
		UserId int64 `binding:"required"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, response.NewErrorResponse(err.Error(), 1))
		return
	}

	if req.UserId < 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("bad user supplied", 2))
		return
	}

	// TODO: Validate userId

	c.Set("UserId", req.UserId)
}
