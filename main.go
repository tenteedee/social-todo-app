package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tenteedee/social-todo-app/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	Init()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	fmt.Print(db)

	now := time.Now().UTC()

	item := types.TodoItem{
		Id:          1,
		Title:       "Buy Milk",
		Images:      "milk.jpg",
		Description: "Buy a gallon of milk from the store",
		Status:      "doing",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	// jsonData, err := json.MarshalIndent(item, "", "  ")

	// if err != nil {
	// 	panic(err)
	// }

	// println(string(jsonData))

	// i := types.TodoItem{}
	// if err := json.Unmarshal([]byte(jsonData), &i); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(i)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("")
			items.GET("/:id")
			items.POST("", CreateItem(db))
			items.PUT("/:id")
			items.DELETE("/:id")
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})

	r.Run(":3000")
}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data types.TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Item created successfully",
			"data":    data.Id,
		})

	}
}
