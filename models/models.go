package models

// Person struct will hold each row of your CSV
type Person struct {
	Name   string `csv:"name"`
	Gender string `csv:"gender"`
}
