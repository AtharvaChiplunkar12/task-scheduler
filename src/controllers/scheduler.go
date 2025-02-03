package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"src/db_files"
	"src/models"

	"github.com/gin-gonic/gin"
)

func ProcessRequest(c *gin.Context, db *sql.DB) {
	var request_dto models.UserRequest_DTO

	if err := c.ShouldBindJSON(&request_dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Invalid Input": err.Error(),
		})
		return
	}
	request, err := convertDTO(request_dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Time format Error": err.Error(),
		})
		return
	}
	id, err := db_files.AddRequestDB(db, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Database Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.RequestResponse{ID: id})
}

func GetStatus(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	fmt.Println(id)
	request, err := db_files.GetRequestDB(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Database Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, request)
}
