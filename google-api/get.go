package google_api

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/fitness/v1"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"
)

// GetDatasetBody
// authorizes with the user's token and aggregate data about weight
func GetDatasetBody(ctx context.Context, startTime, endTime time.Time, dataType string) (*fitness.AggregateResponse, error) {
	flag.Parse()
	config, err := returnConfig([]string{
		fitness.FitnessBodyReadScope,
		fitness.FitnessBodyWriteScope,
	})
	if err != nil {
		log.Println("returnConfig error", err.Error())
		return nil, err
	}

	if *debug {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
			Transport: &logTransport{http.DefaultTransport},
		})
	}

	// returning HTTP client using user's token and configs of the application
	client := authClient(ctx, config)
	svc, err := fitness.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println("NewService error", err.Error())
		return nil, err
	}

	// in AggregateRequest we use milliseconds for StartTimeMillis and EndTimeMillis,
	// while in response we get time in nanoseconds
	payload := fitness.AggregateRequest{
		AggregateBy: []*fitness.AggregateBy{
			{DataTypeName: "com.google."+dataType},
		},
		BucketByTime: &fitness.BucketByTime{
			Period: &fitness.BucketByTimePeriod{
				Type:       "day",
				Value:      1,
				TimeZoneId: "GMT",
			},
		},
		StartTimeMillis: TimeToMillis(startTime),
		EndTimeMillis:   TimeToMillis(endTime),
	}

	weightData, err := svc.Users.Dataset.Aggregate("me", &payload).Do()
	if err != nil {
		return nil, errors.New("Unable to query aggregated weight data:" + err.Error())
	}

	return weightData, nil
}

func GetDatasetHeartRate(ctx context.Context, startTime, endTime time.Time, dataType string) (*fitness.AggregateResponse, error) {
	flag.Parse()
	config, err := returnConfig([]string{
		fitness.FitnessHeartRateWriteScope,
		fitness.FitnessHeartRateReadScope,
	})
	if err != nil {
		log.Println("returnConfig error", err.Error())
		return nil, err
	}

	if *debug {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
			Transport: &logTransport{http.DefaultTransport},
		})
	}

	// returning HTTP client using user's token and configs of the application
	client := authClient(ctx, config)
	svc, err := fitness.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println("NewService error", err.Error())
		return nil, err
	}

	// in AggregateRequest we use milliseconds for StartTimeMillis and EndTimeMillis,
	// while in response we get time in nanoseconds
	payload := fitness.AggregateRequest{
		AggregateBy: []*fitness.AggregateBy{
			{DataTypeName: "com.google."+dataType},
		},
		BucketByTime: &fitness.BucketByTime{
			Period: &fitness.BucketByTimePeriod{
				Type:       "day",
				Value:      1,
				TimeZoneId: "GMT",
			},
		},
		StartTimeMillis: TimeToMillis(startTime),
		EndTimeMillis:   TimeToMillis(endTime),
	}

	weightData, err := svc.Users.Dataset.Aggregate("me", &payload).Do()
	if err != nil {
		return nil, errors.New("Unable to query aggregated weight data:" + err.Error())
	}

	return weightData, nil
}

func GetActivity(ctx context.Context, startTime, endTime time.Time, dataType string) (*fitness.AggregateResponse, error) {
	flag.Parse()
	config, err := returnConfig([]string{
		fitness.FitnessActivityReadScope,
		fitness.FitnessActivityWriteScope,
	})
	if err != nil {
		log.Println("returnConfig error", err.Error())
		return nil, err
	}

	if *debug {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
			Transport: &logTransport{http.DefaultTransport},
		})
	}

	// returning HTTP client using user's token and configs of the application
	client := authClient(ctx, config)
	svc, err := fitness.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println("NewService error", err.Error())
		return nil, err
	}

	// in AggregateRequest we use milliseconds for StartTimeMillis and EndTimeMillis,
	// while in response we get time in nanoseconds
	payload := fitness.AggregateRequest{
		AggregateBy: []*fitness.AggregateBy{
			{DataTypeName: "com.google."+dataType},
		},
		BucketByTime: &fitness.BucketByTime{
			Period: &fitness.BucketByTimePeriod{
				Type:       "day",
				Value:      1,
				TimeZoneId: "GMT",
			},
		},
		StartTimeMillis: TimeToMillis(startTime),
		EndTimeMillis:   TimeToMillis(endTime),
	}

	weightData, err := svc.Users.Dataset.Aggregate("me", &payload).Do()
	if err != nil {
		return nil, errors.New("Unable to query aggregated weight data:" + err.Error())
	}

	return weightData, nil
}


func GetNutritionDataset(ctx context.Context, startTime, endTime time.Time, dataType string) ([]*fitness.Dataset, error)  {
	flag.Parse()

	config, err := returnConfig([]string{
		fitness.FitnessNutritionReadScope,
		fitness.FitnessNutritionWriteScope,
	})
	if err != nil {
		log.Println("returnConfig error", err.Error())
		return nil, err
	}

	if *debug {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
			Transport: &logTransport{http.DefaultTransport},
		})
	}

	// returning HTTP client using user's token and configs of the application
	client := authClient(ctx, config)
	svc, err := fitness.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println("NewService error", err.Error())
		return nil, err
	}

	return NotAggregatedDatasets(svc, startTime, endTime, dataType)
}

func NotAggregatedDatasets(svc *fitness.Service, minTime, maxTime time.Time, dataType string) ([]*fitness.Dataset, error) {
	ds, err := svc.Users.DataSources.List("me").DataTypeName("com.google." + dataType).Do()
	if err != nil {
		log.Println("Unable to retrieve user's data sources:", err)
		return nil, err
	}
	if len(ds.DataSource) == 0 {
		log.Println("You have no data sources to explore.")
		return nil, err
	}

	var dataset []*fitness.Dataset

	// there are different datasources
	// each datasource has datastream
	// we can create our own datastream
	// by default there is datastream called "merged"
	for _, d := range ds.DataSource {
		setID := fmt.Sprintf("%v-%v", minTime.UnixNano(), maxTime.UnixNano())
		data, err := svc.Users.DataSources.Datasets.Get("me", d.DataStreamId, setID).Do()
		if err != nil {
			log.Println("Unable to retrieve dataset:", err.Error())
			return nil, err
		}
		dataset = append(dataset, data)
	}

	return dataset, nil

}