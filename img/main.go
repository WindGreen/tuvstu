package main

import (
	"log"
	"net/http"
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
	router.GET("/:id", GetHandle)
	router.Run(":8082")
}

func GetHandle(c *gin.Context) {
	id := c.Param("id")
	picture, err := tuvstu.FindPicture(id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	log.Printf("%#v\n", picture)
	log.Println(picture.GetLocation())
	c.File(picture.GetLocation())
}
