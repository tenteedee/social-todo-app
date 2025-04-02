package gin_api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tenteedee/social-todo-app/common"
	"github.com/tenteedee/social-todo-app/modules/business"
	"github.com/tenteedee/social-todo-app/modules/item/model"
	"github.com/tenteedee/social-todo-app/modules/item/storage"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		businessLogic := business.NewUpdateItemBusiness(store)

		if err := businessLogic.NewUpdateItemBusiness(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
