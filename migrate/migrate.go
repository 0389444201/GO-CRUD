package main

import (
	"GO-CRUD1/initializers"
	"GO-CRUD1/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(models.Post{})
}
