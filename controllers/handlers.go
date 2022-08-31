package controllers

import (
	"context"
	"devto-google-fit/google-api"
	"devto-google-fit/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetWeight(c *gin.Context) {
	startTime := time.Now().AddDate(0, 0, -31)
	endTime := time.Now()

	dataset, err := google_api.GetDatasetBody(context.Background(), startTime, endTime, models.WeightDataType)
	if err != nil {
		log.Println("GetWeight error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	// parsedData := google_api.ParseData(dataset, models.WeightDataType)
	c.JSON(http.StatusOK, dataset)
	return
}

func GetHeight(c *gin.Context) {
	startTime := time.Now().AddDate(0, 0, -31)
	endTime := time.Now()

	dataset, err := google_api.GetDatasetBody(context.Background(), startTime, endTime, models.HeightDataType)
	if err != nil {
		log.Println("GetHeight error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	parsedData := google_api.ParseData(dataset, models.HeightDataType)
	c.JSON(http.StatusOK, parsedData)
	return
}


func GetHeartRate(c *gin.Context) {
	startTime := time.Now().AddDate(0, 0, -31)
	endTime := time.Now()

	dataset, err := google_api.GetDatasetHeartRate(context.Background(), startTime, endTime, models.HeartRate)
	if err != nil {
		log.Println("GetHeartRate error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	parsedData := google_api.ParseData(dataset, models.HeartRate)
	c.JSON(http.StatusOK, parsedData)
	return
}

func GetHeartPoints(c *gin.Context) {
	startTime := time.Now().AddDate(0, 0, -31)
	endTime := time.Now()

	dataset, err := google_api.GetActivity(context.Background(), startTime, endTime, models.HeartPoints)
	if err != nil {
		log.Println("GetHeartPoints error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	parsedData := google_api.ParseHeartPoints(dataset, models.HeartPoints)
	c.JSON(http.StatusOK, parsedData)
	return
}

func GetActiveMinutes(c *gin.Context) {
	startTime := time.Now().AddDate(0, 0, -31)
	endTime := time.Now()

	dataset, err := google_api.GetActivity(context.Background(), startTime, endTime, models.ActiveMinutes)
	if err != nil {
		log.Println("GetActiveMinutes error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	parsedData := google_api.ParseActivityData(dataset, models.ActiveMinutes)
	c.JSON(http.StatusOK, parsedData)
	return
}

func GetSteps(c *gin.Context) {
	now := time.Now()
	sec := time.Duration(now.Second())
	min := time.Duration(now.Minute())
	hour := time.Duration(now.Hour())

	startTime := now.AddDate(0,0,-31).
		Add(- time.Hour * hour).
		Add(-time.Minute * min).
		Add(-time.Second * sec)

	dataset, err := google_api.GetActivity(context.Background(), startTime, now, models.Steps)
	if err != nil {
		log.Println("GetSteps error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	parsedData := google_api.ParseActivityData(dataset, models.Steps)
	c.JSON(http.StatusOK, parsedData)
	return
}

func GetActivitySegment(c *gin.Context) {
	now := time.Now()
	sec := time.Duration(now.Second())
	min := time.Duration(now.Minute())
	hour := time.Duration(now.Hour())

	startTime := now.AddDate(0,0,-31).
		Add(- time.Hour * hour).
		Add(-time.Minute * min).
		Add(-time.Second * sec)

	dataset, err := google_api.GetActivity(context.Background(), startTime, now, models.ActivitySegment)
	if err != nil {
		log.Println("GetActivitySegment error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	parsedData := google_api.ParseActivitySegment(dataset)
	c.JSON(http.StatusOK, parsedData)
	return
}

func GetCaloriesBurnt(c *gin.Context) {
	now := time.Now()
	sec := time.Duration(now.Second())
	min := time.Duration(now.Minute())
	hour := time.Duration(now.Hour())

	startTime := now.AddDate(0,0,-31).
		Add(- time.Hour * hour).
		Add(-time.Minute * min).
		Add(-time.Second * sec)

	dataset, err := google_api.GetActivity(context.Background(), startTime, now, models.CaloriesBurnt)
	if err != nil {
		log.Println("GetActivitySegment error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	parsedData := google_api.ParseCalories(dataset)
	c.JSON(http.StatusOK, parsedData)
	return
}

func GetHydration(c *gin.Context) {
	now := time.Now()
	sec := time.Duration(now.Second())
	min := time.Duration(now.Minute())
	hour := time.Duration(now.Hour())

	startTime := now.AddDate(0,0,-31).
		Add(- time.Hour * hour).
		Add(-time.Minute * min).
		Add(-time.Second * sec)

	dataset, err := google_api.GetNutritionDataset(context.Background(), startTime, now, models.Hydration)
	if err != nil {
		log.Println("GetHydration error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	parsedData := google_api.ParseHydration(dataset)
	c.JSON(http.StatusOK, parsedData)
	return
}

func GetNutrition(c *gin.Context) {
	now := time.Now()
	sec := time.Duration(now.Second())
	min := time.Duration(now.Minute())
	hour := time.Duration(now.Hour())

	startTime := now.AddDate(0,0,-31).
		Add(- time.Hour * hour).
		Add(-time.Minute * min).
		Add(-time.Second * sec)

	dataset, err := google_api.GetNutritionDataset(context.Background(), startTime, now, models.Nutrition)
	if err != nil {
		log.Println("GetNutrition error:", err.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}

	parsedData := google_api.ParseNutrition(dataset)
	c.JSON(http.StatusOK, parsedData)
	return
}