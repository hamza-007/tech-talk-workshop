package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	uuid "github.com/google/uuid"
	"main/database"
	"main/models"
)

func Add(c *fiber.Ctx) error {
	var book models.Book
	err := c.BodyParser(&book)
	if err != nil {
		return c.SendString("error parsing ")
	}
	book.Id = uuid.New().String()
	stmt, err := database.DB.Prepare("INSERT INTO books (id,title,author,year,edition,price) VALUES (?,?,?,?,?,?) ")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(book.Id, book.Title, book.Author, book.Year, book.Edition, book.Price)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "succes added a new book",
		"book":   book,
	})

}

func GetAllBooks(c *fiber.Ctx) error {
	var books []models.Book
	var book models.Book

	rows, e := database.DB.Query("SELECT * FROM books")

	if e != nil {
		return e
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year, &book.Edition, &book.Price)
		if err != nil {
			return err
		}
		books = append(books, book)
	}
	return c.JSON(fiber.Map{
		"books founded": len(books),
		"books":         books,
	})
}

func GetBook(c *fiber.Ctx) error {
	var book models.Book
	id := c.Params("id")

	row, err := database.DB.Query("SELECT * FROM books WHERE id = ?", id)
	if err != nil {
		return err
	}
	for row.Next() {
		err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Year, &book.Edition, &book.Price)
		if err != nil {
			return err
		}
	}
	if book.Id == "" {
		return c.SendString("book doesn't exist")
	}
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	var book models.Book
	id := c.Params("id")
	err := c.BodyParser(&book)
	if err != nil {
		c.SendString("error parsing data")
	}
	stmt, err := database.DB.Prepare("UPDATE books SET title = ? , author = ? , year = ? , price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(&book.Title, &book.Author, &book.Year, &book.Price, &id)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message ": "succes update",
		"book":     book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	stmt, err := database.DB.Prepare("DELETE FROM books WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message ": "succes delete",
	})
}
