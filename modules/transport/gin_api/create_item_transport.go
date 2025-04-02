package gin_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tenteedee/social-todo-app/common"
	"github.com/tenteedee/social-todo-app/modules/business"
	"github.com/tenteedee/social-todo-app/modules/item/model"
	"github.com/tenteedee/social-todo-app/modules/item/storage"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		businessLogic := business.NewCreateItemBusiness(store)

		if err := businessLogic.CreateItemBusiness(c.Request.Context(), data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
