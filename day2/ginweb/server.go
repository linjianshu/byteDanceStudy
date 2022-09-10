package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"src/ByteDance/day2/ginweb/controller"
	"src/ByteDance/day2/ginweb/repository"
)

func main() {
	if err := Init("E:\\project\\GOproject\\src\\ByteDance\\day2\\ginweb\\data\\"); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(context *gin.Context) {
		topicId := context.Param("id")
		data := controller.QueryPageInfo(topicId)
		context.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
