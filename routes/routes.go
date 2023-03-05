package routes

import(
	"github.com/gofiber/fiber/v2"
	"main/controllers"
	
)


func Setup(app *fiber.App){
	//add new book
	app.Post("/add",controllers.Add)

	//get all books
	app.Get("/book/all",controllers.GetAllBooks)

	//get book by id
	app.Get("book/:id",controllers.GetBook)

	//update existing book by id 
	app.Put("/update/:id",controllers.UpdateBook)
	
	//delete a existing book by id 
	app.Delete("/delete/:id",controllers.DeleteBook)

}