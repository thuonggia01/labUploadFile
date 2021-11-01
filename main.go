package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
type image struct {
	url string
}

var img = image{url: ""}

func index(c *gin.Context) {
	if img.url != "" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"url": img.url,
		})
		fmt.Println(img.url)
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func uploadfile(c *gin.Context) {
	file, err := c.FormFile("image")
	path := fmt.Sprintf("./assets/upload/%v", file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	err = c.SaveUploadedFile(file, "./assets/upload/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	img.url = path
	//http.Redirect(c.Writer, c.Request, "/", 302)
	index(c)
}
func main()  {
	router := gin.Default()
	router.Static("/assets","./assets")
	router.LoadHTMLGlob("./assets/**/*.html")
	// r.MaxMultipartMemory = 8 << 18 //2mib
	router.GET("/", index)
	router.POST("/upload", uploadfile)
	router.Run()
}