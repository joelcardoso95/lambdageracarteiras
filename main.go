package main

import (
	"bytes"
	"context"
	"log"
	"text/template"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lambdageracarteiras/models"
	"github.com/lambdageracarteiras/service"
)

// RequestPDF struct
type RequestPDF struct {
	body []byte
}

func main() {
	lambda.Start(LambdaExecute)
}

func LambdaExecute(context context.Context) (string, error) {
	s3File, err := service.DownloadFileFromS3Bucket("carteiras-adviladiva", "excel/users.csv")
	if err != nil {
		log.Fatalf("Failed to download from S3: %v", err)
	}
	people, err := service.ReadCSV(s3File.Name())
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
		return "", err
	}

	htmlFile, err := service.DownloadHtmlFromS3Bucket("carteiras-adviladiva", "html/carteira.html")
	if err != nil {
		log.Fatalf("Failed to download HTML from S3: %v", err)

	}

	for _, person := range people {
		pdfRequest := &RequestPDF{}
		if err != nil {
			log.Fatalf("Failed to read HTML file: %v", err)
			return "", err
		}
		err = pdfRequest.ParseTemplate(htmlFile.Name(), *person)
		if err != nil {
			log.Fatalf("Failed to parse template: %v", err)
			return "", err
		}
		log.Println("Successfully Generated HTML Files Bytes", pdfRequest.body)
		service.UploadFileToS3Bucket("carteiras-adviladiva", "pdf/"+person.Name+".html", pdfRequest.body)
		log.Println("Successfully Uploaded Files to S3 Bucket")
	}

	return "Successfully generated HTML Files", nil
}

// write the code to parse template and return de buffer
func (r *RequestPDF) ParseTemplate(htmlTemplate string, person models.Person) error {
	t, err := template.ParseFiles(htmlTemplate)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, person); err != nil {
		return err
	}
	r.body = buf.Bytes()
	return nil
}
