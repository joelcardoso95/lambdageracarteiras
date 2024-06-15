package service

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/lambdageracarteiras/models"
)

func ReadCSV(filePath string) ([]*models.Person, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var people []*models.Person
	if err := gocsv.UnmarshalFile(file, &people); err != nil {
		return nil, err
	}

	return people, nil
}
