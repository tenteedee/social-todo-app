package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiFeature(c *gin.Context, db *gorm.DB) *gorm.DB {
	// Filters
	for key, values := range c.Request.URL.Query() {
		if key != "sort" && key != "page" && key != "limit" {
			// Apply filter to the query
			db = db.Where(key+" = ?", values[0])
		}
	}

	// Get the query parameters from the request
	query := c.Request.URL.Query()

	// Get the page and limit parameters
	page := query.Get("page")
	limit := query.Get("limit")

	// Set default values if not provided
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	// Convert page and limit to integers
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Calculate offset for pagination
	offset := (pageInt - 1) * limitInt

	// Apply pagination to the query
	db = db.Offset(offset).Limit(limitInt)

	return db
}
