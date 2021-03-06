package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func main() {
	svc := cloudwatch.New(session.New())

	params := &cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{ // Required
			{ // Required
				MetricName: aws.String("MetricName"), // Required
				Dimensions: []*cloudwatch.Dimension{
					{ // Required
						Name:  aws.String("DimensionName"),  // Required
						Value: aws.String("DimensionValue"), // Required
					},
					// More values...
				},
				StatisticValues: &cloudwatch.StatisticSet{
					Maximum:     aws.Float64(1.0), // Required
					Minimum:     aws.Float64(1.0), // Required
					SampleCount: aws.Float64(1.0), // Required
					Sum:         aws.Float64(1.0), // Required
				},
				Timestamp: aws.Time(time.Now()),
				Unit:      aws.String("StandardUnit"),
				Value:     aws.Float64(1.0),
			},
			// More values...
		},
		Namespace: aws.String("Namespace"), // Required
	}
	resp, err := svc.PutMetricData(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
