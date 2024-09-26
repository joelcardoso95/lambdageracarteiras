package main

import (
	"bytes"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lambdageracarteiras/models"
	"github.com/lambdageracarteiras/service"
	"log"
	"text/template"
	"time"
)

// RequestPDF struct
type RequestPDF struct {
	body []byte
}

type PageData struct {
	People []models.Person
}

func main() {
	lambda.Start(LambdaExecute)
}

func LambdaExecute(context context.Context) (string, error) {
	s3File, err := service.DownloadFileFromS3Bucket("carteiras-adviladiva", "excel/cadastro-igreja.csv")
	if err != nil {
		log.Fatalf("Failed to download from S3: %v", err)
	}
	peopleCSV, err := service.ReadCSV(s3File.Name())
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
		return "", err
	}

	htmlFile, err := service.DownloadHtmlFromS3Bucket("carteiras-adviladiva", "html/index.html")
	if err != nil {
		log.Fatalf("Failed to download HTML from S3: %v", err)

	}

	var persons []models.Person
	for _, person := range peopleCSV {
		if person.Name != "" {
			birthday, err := time.Parse("1/2/2006", person.BirthDay)
			if err != nil {
				log.Fatalf("Failed to parse date: %v", err)
				return "", err
			}
			persons = append(persons, models.Person{
				Name:     person.Name,
				BirthDay: birthday.Format("02/01/2006"),
				CPF:      person.CPF,
			})
		}
	}

	pageData := PageData{
		People: persons,
	}

	pdfRequest := &RequestPDF{}
	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
		return "", err
	}
	err = pdfRequest.ParseTemplate(htmlFile.Name(), pageData)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
		return "", err
	}
	log.Println("Successfully Generated HTML Files Bytes", pdfRequest.body)
	service.UploadFileToS3Bucket("carteiras-adviladiva", "pdf/"+"carteiras-geradas"+".html", pdfRequest.body)
	log.Println("Successfully Uploaded Files to S3 Bucket")

	return "Successfully generated HTML Files", nil
}

// write the code to parse template and return de buffer
func (r *RequestPDF) ParseTemplate(htmlTemplate string, pageData PageData) error {
	t, err := template.ParseFiles(htmlTemplate)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, pageData); err != nil {
		return err
	}
	r.body = buf.Bytes()
	return nil
}
