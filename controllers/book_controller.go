package controllers

import (
	"github.com/Ilham-muttaqien17/learn-restful-go/dto"
	"github.com/Ilham-muttaqien17/learn-restful-go/services"
	"github.com/Ilham-muttaqien17/learn-restful-go/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookController struct {
	bookService *services.BookService
}

func RegisterBookController() *BookController {
	return &BookController{
		bookService: &services.BookService{},
	}
}


func (c *BookController) Index(ctx *fiber.Ctx) error {
	pagination := utils.BuildPaginationParams(ctx)

	response := c.bookService.GetAllBooks(&pagination)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *BookController) Show(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id")

	book, err := c.bookService.GetDetailBook(bookId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "Data not found")
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"data": book,
	})
}

func (c *BookController) Create(ctx *fiber.Ctx) error {
	var bookDTO dto.BookDTO
	
	// Parse data from request body
	if err := ctx.BodyParser(&bookDTO); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Validate request
	parsedBook, errValidation := utils.Validator(bookDTO);
	if  errValidation != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(map[string]interface{} {
			"errors": errValidation,
		})
	}

	savedBook, err := c.bookService.CreateBook(parsedBook)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created succesfully",
		"data": savedBook,
	})
}

func (c *BookController) Update(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id")
	var bookDTO dto.BookDTO

	// Parse data from request body
	if err := ctx.BodyParser(&bookDTO); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Validate request
	parsedBook, errValidation := utils.Validator(bookDTO);
	if  errValidation != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(map[string]interface{} {
			"errors": errValidation,
		})
	}

	savedBook, err := c.bookService.UpdateBook(parsedBook, bookId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "Data not found")
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book updated succesfully",
		"data": savedBook,
	})
}

func (c *BookController) Destroy(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id")

	if err := c.bookService.DeleteBook(bookId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "Data not found")
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted succesfully",
	})
}