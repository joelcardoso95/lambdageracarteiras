package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
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
		err := template.Must(template.New(person.Name).Parse("templates/carteira.html")).Execute(&htmlBuffer, person)
		if err != nil {
			log.Fatalf("Failed to execute template: %v", err)
			return "", err
		}
		pdfGen, err := pdf.NewPDFGenerator()
		if err != nil {
			log.Fatalf("Failed to create PDF generator: %v", err)
			return "", err
		}
		page := pdf.NewPageReader(bytes.NewReader(htmlBuffer.Bytes()))
		pdfGen.AddPage(page)
		pdfGen.WriteFile("/tmp/pdf/" + person.Name + ".pdf")
	}

	return "Successfully generated PDF Files", nil
}
