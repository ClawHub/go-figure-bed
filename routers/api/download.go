package api

import (
	"github.com/gin-gonic/gin"
	"go-figure-bed/service"
	"net/http"
)

//获取图片
func GetImage(c *gin.Context) {
	mainUrl := c.Param("mainUrl")
	url := service.Download(mainUrl, "guest")
	if url != "" {
		c.Redirect(http.StatusMovedPermanently, url)
	}
	c.String(http.StatusNotFound, "image not found")
}
