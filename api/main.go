package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"tuvstu"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := tuvstu.InitMgo("localhost")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/pictures", UploadHandle)
	router.Run(":8081")
}

func UploadHandle(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	files := form.File["files"]
	now := time.Now()
	path := fmt.Sprintf("../%s/%d/%02d/%02d", "pictures", now.Year(), now.Month(), now.Day())
	for _, file := range files {
		picture := tuvstu.NewPicture(file.Filename, path)
		if err := c.SaveUploadedFile(file, picture.GetLocation()); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
			return
		}
		picture.Save()
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}
