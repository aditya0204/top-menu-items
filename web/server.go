package web

import (
	"log"
	"net/http"
	"top-menu-items/models"
	"top-menu-items/services"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/upload-menu-log", UploadMenuLogController)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func UploadMenuLogController(ctx *gin.Context) {
	var req models.UploadMenuLogRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK,models.Response{
			StatusCode: http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}
	data,err:=services.UploadMenuLogFileService(req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK,models.Response{
			StatusCode: http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,models.Response{
		StatusCode: http.StatusInternalServerError,
		Error: "",
		Data: data,
	})

}
