package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func main() {
	svc := cloudwatch.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

	params := &cloudwatch.DescribeAlarmHistoryInput{
		AlarmName:       aws.String("AlarmName"),
		EndDate:         aws.Time(time.Now()),
		HistoryItemType: aws.String("HistoryItemType"),
		MaxRecords:      aws.Int64(1),
		NextToken:       aws.String("NextToken"),
		StartDate:       aws.Time(time.Now()),
	}

	resp, err := svc.DescribeAlarmHistory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the code and
		// Message from an error
		fmt.Println(err.Error())
		return
	}

	// Prety-print the response data.
	fmt.Println(resp)

}
