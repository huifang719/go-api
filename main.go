package main

import (
	"example/TEST-SERVER-GO/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Type, Authorization, Accept, Accept-Language, X-Requested-With, XMLHttpRequest")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Accept-Language, X-Requested-With, XMLHttpRequest")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")	
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}	

func getGauge(c *gin.Context){
	// queryParams := c.Request.URL.Query()
	appID := c.Query("appID")
	objectID := c.Query("objectID")

	if appID == "" || objectID == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Please enter App ID or/and Object ID"})
		return
	}
	if appID != "da8c91fe-88fd-45ef-b573-596621f7ec6f" || objectID != "576c5fc2-e038-4c90-b682-4e17b2fd846e" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Incorrect App ID or/and Object ID"})
		return
	}

	gauge := controller.CreateGauge()
	c.IndentedJSON(http.StatusOK, gauge)
}	

func main() {	
	router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/api/qlik/object", getGauge)
	router.Run(":8080")
}



