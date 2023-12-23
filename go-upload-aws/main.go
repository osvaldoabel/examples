package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	CONTENT_TYPE = "multipart/form-data"
)

var (
	AWS_REGION      = ""
	AWS_ACCESS_KEY  = ""
	AWS_SECRET_KEY  = ""
	AWS_BUCKET_NAME = ""
)

func getFile() ([]byte, error) {
	return os.ReadFile("./tmp/file.text")
}

func main() {
	file, err := getFile()
	if err != nil {
		panic("could not read file ")
	}

	s3Config := &aws.Config{
		Region:      aws.String(AWS_REGION),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_KEY, ""),
	}

	s3Session := session.New(s3Config)
	uploader := s3manager.NewUploader(s3Session)

	destinationName := "new-filename.txt"

	input := &s3manager.UploadInput{
		Bucket:      aws.String(AWS_BUCKET_NAME), // bucket's name
		Key:         aws.String(destinationName), // files destination location
		Body:        bytes.NewReader(file),       // content of the file
		ContentType: aws.String(CONTENT_TYPE),    // content type
	}
	output, err := uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		log.Default().Fatal(err)
	}

	fmt.Println("done ---> =D", output)
}
