package models

// Person struct will hold each row of your CSV
type Person struct {
	Name     string `csv:"Nome"`
	BirthDay string `csv:"Data Nascimento"`
	Position string `csv:"Data Nascimento"`
	CPF      string `csv:"Selecione o Cargo"`
}
