package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		diskUUID := c.PostForm("diskUUID")
		log.Println("Upload Requested for Disk : ", diskUUID)

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		targetDir := "./temp/"
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			os.Mkdir(targetDir, os.ModePerm)
		}

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, targetDir+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		log.Println("Upload success - fileName : ", filename)

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully for diskUUID=%s.", file.Filename, diskUUID))
	})
	router.Run(":8080")
}
