package service

import (
	"bytes"
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
	file, err := os.Create("/tmp/downloaded.csv")
	if err != nil {
		return nil, err
	}

	log.Printf("starting s3 download")
	// download the file from S3
	n, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &dowloadFileName,
	})
	if err != nil {
		log.Fatalf("Failed to download from S3: %v", err)
	}

	log.Printf("file downloaded, %d bytes\n", n)
	return file, nil
}

func DownloadHtmlFromS3Bucket(bucket string, dowloadFileName string) (*os.File, error) {
	// create a session
	s3Session := session.Must(session.NewSession())

	// create a downloader
	downloader := s3manager.NewDownloader(s3Session)

	// create a file to write the S3 Object contents to.
	file, err := os.Create("/tmp/downloaded.html")
	if err != nil {
		return nil, err
	}

	log.Printf("starting s3 download")
	// download the file from S3
	n, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &dowloadFileName,
	})
	if err != nil {
		log.Fatalf("Failed to download from S3: %v", err)
	}

	log.Printf("file downloaded, %d bytes\n", n)
	return file, nil
}

func UploadFileToS3Bucket(bucket string, uploadFileName string, file []byte) error {
	log.Println("Uploaling S3 File", uploadFileName)
	// create a session
	s3Session := session.Must(session.NewSession())

	// create an uploader
	uploader := s3manager.NewUploader(s3Session)

	// upload the file to S3
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &uploadFileName,
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		log.Fatalf("Failed to upload file to S3: %v", err)
		return err
	}

	log.Printf("file uploaded successfully")
	return nil
}
