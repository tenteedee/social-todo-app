package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	ItemController "github.com/tenteedee/social-todo-app/controllers"
)

func main() {
	Init()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	// fmt.Print(db)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", ItemController.GetItems(db))
			items.GET("/:id", ItemController.GetItemById(db))
			items.POST("", ItemController.CreateItem(db))
			items.PUT("/:id", ItemController.UpdateItem(db))
			items.DELETE("/:id", ItemController.DeleteItem(db))
		}
	}

	r.Run(":3000")
}
