package google_api

import (
	"devto-google-fit/models"
	"fmt"
	"google.golang.org/api/fitness/v1"
	"strconv"
)

func ParseData(ds *fitness.AggregateResponse, dataType string) []models.Measurement {
	var data []models.Measurement

	for _, res := range ds.Bucket {
		for _, ds := range res.Dataset {
			for _, p := range ds.Point {
				var row models.Measurement
				row.AvValue = p.Value[0].FpVal
				row.MinValue = p.Value[1].FpVal
				row.MaxValue = p.Value[2].FpVal
				row.StartTime = NanosToTime(p.StartTimeNanos)
				row.EndTime = NanosToTime(p.EndTimeNanos)
				row.Type = dataType
				data = append(data, row)
			}
		}
	}
	return data
}

func ParseHeartPoints(ds *fitness.AggregateResponse, dataType string) []models.Measurement {
	var data []models.Measurement

	for _, res := range ds.Bucket {
		for _, ds := range res.Dataset {
			for _, p := range ds.Point {
				var row models.Measurement
				row.Value = p.Value[0].FpVal
				row.StartTime = NanosToTime(p.StartTimeNanos)
				row.EndTime = NanosToTime(p.EndTimeNanos)
				row.Type = dataType
				data = append(data, row)
			}
		}
	}
	return data
}

func ParseActivityData(ds *fitness.AggregateResponse, dataType string) []models.ActiveMinute {
	// strange that in documentation it is stated that response is in milliseconds but actually in minutes
	var data []models.ActiveMinute

	for _, res := range ds.Bucket {
		for _, ds := range res.Dataset {
			for _, p := range ds.Point {
				var row models.ActiveMinute
				row.Value = p.Value[0].IntVal
				row.StartTime = NanosToTime(p.StartTimeNanos)
				row.EndTime = NanosToTime(p.EndTimeNanos)
				row.Type = dataType
				data = append(data, row)
			}
		}
	}
	return data
}

func ParseActivitySegment(ds *fitness.AggregateResponse) []models.ActivitySegments {
	var data []models.ActivitySegments

	for _, res := range ds.Bucket {
		for _, ds := range res.Dataset {
			for _, p := range ds.Point {
				var row models.ActivitySegments
				row.ActivityType = models.ActivitySegmentMap[p.Value[0].IntVal]
				row.Minutes = float64(p.Value[1].IntVal)/60000
				row.SessionNum = p.Value[2].IntVal
				row.StartTime = NanosToTime(p.StartTimeNanos)
				row.EndTime = NanosToTime(p.EndTimeNanos)
				data = append(data, row)
			}
		}
	}
	return data
}

func ParseCalories(ds *fitness.AggregateResponse) []models.Calories {
	var data []models.Calories

	for _, res := range ds.Bucket {
		for _, ds := range res.Dataset {
			for _, p := range ds.Point {
				var row models.Calories
				row.Amount = p.Value[0].FpVal
				row.StartTime = NanosToTime(p.StartTimeNanos)
				row.EndTime = NanosToTime(p.EndTimeNanos)
				data = append(data, row)
			}
		}
	}
	return data
}

func ParseHydration(datasets []*fitness.Dataset) []models.HydrationStruct {
	var data []models.HydrationStruct

	for _, ds := range datasets {
		var value float64
		for _, p := range ds.Point {
			for _, v := range p.Value {
				valueString := fmt.Sprintf("%.3f", v.FpVal)
				value, _ = strconv.ParseFloat(valueString, 64)
			}
			var row models.HydrationStruct
			row.StartTime = NanosToTime(p.StartTimeNanos)
			row.EndTime = NanosToTime(p.EndTimeNanos)
			// liters to milliliters
			row.Amount = int(value * 1000)
			data = append(data, row)
		}
	}
	return data
}

func ParseNutrition(datasets []*fitness.Dataset) []models.NutritionStruct {
	var data []models.NutritionStruct

	for _, ds := range datasets {
		for _, p := range ds.Point {
			var row models.NutritionStruct
			for _, mapVal := range p.Value[0].MapVal {
				// there we can get more data (s.t. fat, carbs, protein, etc.) if it exists
				if mapVal.Key == models.NutrientCalories {
					row.Calories = int(mapVal.Value.FpVal)
				}
			}
			row.Type = models.MealType[int(p.Value[1].IntVal)]
			row.Name = p.Value[2].StringVal
			row.StartTime = NanosToTime(p.StartTimeNanos)
			row.EndTime = NanosToTime(p.EndTimeNanos)

			data = append(data, row)
		}
	}
	return data
}