package service

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/lambdageracarteiras/models"
)

func ReadCSV(filePath string) ([]*models.Person, error) {
	log.Printf("starting csv read")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var people []*models.Person
	if err := gocsv.UnmarshalFile(file, &people); err != nil {
		return nil, err
	}

	log.Printf("read %d people from csv\n", len(people))
	return people, nil
}
