package service

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadFileFromS3Bucket(bucket string, dowloadFileName string) (*os.File, error) {
	// create a session
	s3Session := session.Must(session.NewSession())

	// create a downloader
	downloader := s3manager.NewDownloader(s3Session)

	// create a file to write the S3 Object contents to.
	file, err := os.Create("downloaded.csv")
	if err != nil {
		return nil, err
	}

	// download the file from S3
	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &dowloadFileName,
	})
	if err != nil {
		log.Fatalf("Failed to download CSV: %v", err)
	}

	return file, nil
}
