package main

import (
	"github.com/GorrillaProcess/go-crud/initializers"
	"github.com/GorrillaProcess/go-crud/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
