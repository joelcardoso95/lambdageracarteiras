package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lambdageracarteiras/service"
)

func main() {
	lambda.Start(LambdaExecute)
}

func LambdaExecute(context context.Context) (string, error) {
	service.DownloadFileFromS3Bucket("carteiras-adviladiva/excel", "users.csv")
	//people, err := service.ReadCSV(file.Name())
	/*if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
		return "", err
	}

	for _, person := range people {
		fmt.Printf("Name: %s, Gender: %s\n", person.Name, person.Gender)
	}

	return "Successfully processed CSV", nil*/
}
