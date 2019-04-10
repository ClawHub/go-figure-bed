package api

import (
	"github.com/gin-gonic/gin"
	"go-figure-bed/pkg/app"
	"go-figure-bed/pkg/e"
	"go-figure-bed/service"
	"net/http"
)

//检测并修复库中失效图片url
func AutoCheckRepair(c *gin.Context) {
	appG := app.Gin{C: c}
	//调用service处理
	service.AutoCheckRepair()
	appG.Response(http.StatusOK, e.SUCCESS, "Auto Check Repair Finish")
}
