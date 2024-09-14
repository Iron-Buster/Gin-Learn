package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	// 上传单个文件
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	// 上传多个文件
	r.POST("/upload2", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})
}
