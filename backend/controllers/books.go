package controllers

import (
	"github.com/gin-gonic/gin"
	"backend/models"
	"net/http"
)

func FindBooks(c *gin.Context) {
	var books = []models.Book{
		{ID: 1, Title: "A Tale of Two Citiess", Author: "Charles Dickens"},
		{ID: 2, Title: "A Tale of Two Cities", Author: "Charles Dickens"},
		{ID: 3, Title: "A Tale of Two Cities", Author: "Charles Dickens"},
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}
