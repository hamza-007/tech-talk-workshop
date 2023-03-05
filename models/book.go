package models


type Book struct{
	Id 			string  `json:"id"`
	Title 		string 	`json:"title"`
	Author		string	`json:"author"`
	Year 		uint64 	`json:"year"`
	Edition		string	`json:"edition"`
	Price		float32	`json:"price"`					
}