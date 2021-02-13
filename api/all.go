package api

import (
	"github.com/gin-gonic/gin"
	v1 "gitlab.devgroup.tech/shkolkovo/romanych/api/v1"
)

func AllApi(r *gin.Engine) {
	v1.AllApi(r)
}
