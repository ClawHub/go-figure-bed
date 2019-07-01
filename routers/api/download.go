package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-figure-bed/service"
	"net/http"
)

//获取图片
func GetImage(c *gin.Context) {
	mainUrl := c.Param("mainUrl")
	email := c.DefaultQuery("email", "guest")
	fmt.Println(mainUrl)
	fmt.Println(email)
	url := service.Download(mainUrl, email)
	if url != "" {
		c.Redirect(http.StatusMovedPermanently, url)
	}
	c.String(http.StatusNotFound, "image not found")
}
