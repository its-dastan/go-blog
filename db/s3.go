package db

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"github.com/joho/godotenv"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var s3Session *s3.S3

const (
	bucketName = "aws_bucket_name"
	region     = "aws_bucket_region"
	accessId   = "aws_access_id"
	accessKey  = "aws_secret_key"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s3Session = s3.New(session.Must(session.NewSession(&aws.Config{Region: aws.String(os.Getenv(region)), Credentials: credentials.NewStaticCredentials(os.Getenv(accessId), os.Getenv(accessKey), "")})))
}

func ListBucket() (resp *s3.ListBucketsOutput) {
	resp, err := s3Session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}
	return resp
}

// UploadFileToS3 saves a file to aws bucket and returns the url to // the file and an error if there's any
func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	// create a unique file name for the file
	tempFileName := "pictures/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3Session.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("go-blog-2022"),
		Key:    aws.String(tempFileName),
		//ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	return tempFileName, err
}

func GetImage(key string) []byte {
	params := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv(bucketName)), // Required
		Key:    aws.String(key),                   // Required
	}
	resp, err := s3Session.GetObject(params)
	if err != nil {
		fmt.Println(err)
	}
	buffer := make([]byte, *resp.ContentLength)
	defer resp.Body.Close()
	var bbuffer bytes.Buffer
	for true {

		num, rerr := resp.Body.Read(buffer)
		if num > 0 {
			bbuffer.Write(buffer[:num])
		} else if rerr == io.EOF || rerr != nil {
			break
		}
	}
	return buffer
}
