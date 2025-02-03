package routing

import (
	"database/sql"
	"src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.POST("/schedule", func(c *gin.Context) {
		controllers.ProcessRequest(c, db)
	})
	
	// Define the GET route with an ID parameter
	router.GET("/status/:id", func(c *gin.Context) {
		controllers.GetStatus(c, db)
	})

	return router
}
