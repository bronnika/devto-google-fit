package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func RunAllRoutes() {
	r := gin.Default()
	f, err := os.Create("logs")

	if err != nil {
		fmt.Println("file create error", err.Error())
	}

	log.SetOutput(f)
	log.SetOutput(gin.DefaultWriter)

	r.GET("/weight", GetWeight)
	r.GET("/height", GetHeight)
	r.GET("/heart_rate", GetHeartRate)
	r.GET("/heart_points", GetHeartPoints)
	r.GET("/active_minutes", GetActiveMinutes)
	r.GET("/steps", GetSteps)
	r.GET("/activity_segment", GetActivitySegment)
	r.GET("/calories_burnt", GetCaloriesBurnt)
	r.GET("/hydration", GetHydration)
	r.GET("/nutrition", GetNutrition)
	_ = r.Run(":8080")
}

