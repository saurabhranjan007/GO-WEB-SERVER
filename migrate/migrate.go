package main

import (
	initializers "github.com/saurabhranjan007/go-server/Initializers"
	models "github.com/saurabhranjan007/go-server/Models"
)

func init() {
	initializers.LoadEnvvariables()
	initializers.DatabaseConnection()
}

// MIGRATING DB: required params; DB Connection Pointer and Post Model Pointer

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
