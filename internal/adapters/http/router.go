package http

import "github.com/gin-gonic/gin"

func NewRouter(env string) *gin.Engine {
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		return gin.New()
	}
	return gin.Default()
}
