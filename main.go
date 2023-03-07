package main

import (
	"GO-CRUD1/controllers"
	"GO-CRUD1/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/", controllers.PostCreate)
	r.PUT("/up/:id", controllers.UpdateDB)
	r.GET("/show", controllers.ShowAll)
	r.GET("/show/:id", controllers.ShowByID)
	r.Run()
}
