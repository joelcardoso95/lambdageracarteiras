package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

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
	}

	return "Successfully generated HTML Files", nil
}
