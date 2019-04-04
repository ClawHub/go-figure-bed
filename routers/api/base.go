package api

import (
	"github.com/gin-gonic/gin"
	"go-figure-bed/pkg/app"
	"go-figure-bed/pkg/e"
	"net/http"
)

//健康检查
func Welcome(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, "welcome go-figure-bed")
}
