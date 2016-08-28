package main

import (
	"fmt"
	"io/ioutil"

	"github.com/kataras/iris"
	"github.com/spf13/viper"
	"gopkg.in/h2non/bimg.v0"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	iris.Post("/resize", func(ctx *iris.Context) {
		info, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, iris.Map{"error": err.Error() + " : 'image'"})
			return
		}

		file, err := info.Open()
		defer file.Close()
		_ = info.Filename

		buffer, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Error creating in memory buffer")
			return
		}

		image, err := bimg.NewImage(buffer).Resize(300, 300)
		if err != nil {
			fmt.Println("Error resizing the image")
			return
		}

		ctx.Data(iris.StatusOK, image)
	})

	port := viper.GetString("Server.Port")
	iris.Listen(":" + port)
}
