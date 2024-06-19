package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lambdageracarteiras/service"
)

func main() {
	lambda.Start(LambdaExecute)
}

func LambdaExecute(context context.Context) (string, error) {
	var htmlBuffer bytes.Buffer
	s3File, err := service.DownloadFileFromS3Bucket("carteiras-adviladiva", "excel/users.csv")
	if err != nil {
		log.Fatalf("Failed to download from S3: %v", err)
	}
	people, err := service.ReadCSV(s3File.Name())
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
		return "", err
	}

	for _, person := range people {
		fmt.Printf("Name: %s, Gender: %s\n", person.Name, person.Gender)

		tmpl, err := template.ParseFiles("carteira.html")
		if err != nil {
			log.Fatalf("Failed to execute template: %v", err)
			return "", err
		}
		err = tmpl.Execute(&htmlBuffer, person)
		log.Println("Successfully generated HTML Files")
		if err != nil {
			log.Fatalf("Failed to execute template: %v", err)
			return "", err
		}
		pdfGen, err := wkhtmltopdf.NewPDFGenerator()
		if err != nil {
			log.Fatalf("Failed to create PDF generator: %v", err)
			return "", err
		}
		page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlBuffer.Bytes()))
		pdfGen.AddPage(page)
		pdfGen.Orientation.Set(wkhtmltopdf.OrientationPortrait)
		pdfGen.WriteFile("/tmp/pdf/" + person.Name + ".pdf")
		service.UploadFileToS3Bucket("carteiras-adviladiva", "pdf/"+person.Name+".pdf", pdfGen.Bytes())
		log.Println("Successfully Uploaded PDF Files to S3 Bucket")
	}

	return "Successfully generated PDF Files", nil
}
