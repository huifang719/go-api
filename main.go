package main

import (
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
func createGauge() (gauge) {
	gauge := gauge{
		VisType:    "gauge",
		Dimensions: []string{},
		Measurements: []string{
			"% of Patients Admitted to Hospital Emergency Department",
		},
		MeasureAxis: struct {
			Min int `json:"min"`
			Max int `json:"max"`
		}{
			Min: 0,
			Max: 7935345,
		},
		SegmentInfo: struct {
			Limits []struct {
				Value    int  `json:"value"`
				Gradient bool `json:"gradient"`
			} `json:"limits"`
			PaletteColors []struct {
				Color string `json:"color"`
				Index int    `json:"index"`
			} `json:"paletteColors"`
			Colors []struct {
				Color int `json:"color"`
			} `json:"colors"`
		}{
			Limits: []struct {
				Value    int  `json:"value"`
				Gradient bool `json:"gradient"`
			}{
				{
					Value:    2645115,
					Gradient: false,
				},
				{
					Value:    5290230,
					Gradient: false,
				},
			},
			PaletteColors: []struct {
				Color string `json:"color"`
				Index int    `json:"index"`
			}{
				{
					Color: "#83af9b",
					Index: 3,
				},
				{
					Color: "#10cfc9",
					Index: 5,
				},
				{
					Color: "#006580",
					Index: 6,
				},
			},
			Colors: []struct {
				Color int `json:"color"`
			}{
				{
					Color: 3,
				},
			},
		},
		ExactData: [][]int{
			{
				2440522,
			},
		},
	}
	return gauge
}

type gauge struct {
	VisType      string   `json:"visType"`
	Dimensions   []string `json:"dimensions"`
	Measurements []string `json:"measurements"`
	MeasureAxis  struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"measureAxis"`
	SegmentInfo struct {
		Limits []struct {
			Value    int  `json:"value"`
			Gradient bool `json:"gradient"`
		} `json:"limits"`
		PaletteColors []struct {
			Color string `json:"color"`
			Index int    `json:"index"`
		} `json:"paletteColors"`
		Colors []struct {
			Color int `json:"color"`
		} `json:"colors"`
	} `json:"segmentInfo"`
	ExactData [][]int `json:"exactData"`
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

	gauge := createGauge()
	c.IndentedJSON(http.StatusOK, gauge)
}	

func main() {	
	router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/api/qlik/object", getGauge)
	router.Run(":8080")
}



