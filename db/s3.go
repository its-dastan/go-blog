package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var (
	s3Session  *s3.S3
	bucketName = os.Getenv("aws_bucket_name")
	region     = os.Getenv("aws_bucket_region")
	accessId   = os.Getenv("aws_access_id")
	accessKey  = os.Getenv("aws_secret_key")
)

func init() {
	fmt.Println(region)
	s3Session = s3.New(session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-south-1")})))
}

func ListBucket() (resp *s3.ListBucketsOutput) {
	resp, err := s3Session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}
	return resp
}
