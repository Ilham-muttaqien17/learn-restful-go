package services

import (
	"github.com/Ilham-muttaqien17/learn-restful-go/dto"
	"github.com/Ilham-muttaqien17/learn-restful-go/models"
)

type BookService struct{}

func (c *BookService) GetAllBooks() []models.Book {
	var books []models.Book

	models.DB.Find(&books)

	return books
}

func (c *BookService) GetDetailBook(id string) (models.Book, error) {

	var book models.Book;

	if err := models.DB.First(&book, id).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (c *BookService) CreateBook(bookDTO dto.BookDTO) (models.Book, error) {
	var savedBook models.Book

	savedBook.Title = bookDTO.Title
	savedBook.Author = bookDTO.Author
	savedBook.Description = bookDTO.Description
	savedBook.PublishDate = bookDTO.PublishDate
	savedBook.Tags = bookDTO.Tags
	savedBook.Detail = &models.Detail{
		Price: bookDTO.Detail.Price,
		Weight: bookDTO.Detail.Weight,
	}


	if err := models.DB.Create(&savedBook).Error; err != nil {
		return savedBook, err
	}

	return savedBook, nil
}

func (c *BookService) UpdateBook(bookDTO dto.BookDTO, id string) (models.Book, error) {
	var savedBook models.Book

	if err := models.DB.First(&savedBook, id).Error; err != nil {
		return savedBook, err
	}

	savedBook.Title = bookDTO.Title
	savedBook.Description = bookDTO.Description
	savedBook.Author = bookDTO.Author
	savedBook.PublishDate = bookDTO.PublishDate
	savedBook.Tags = bookDTO.Tags
	savedBook.Detail = &models.Detail{
		Price: bookDTO.Detail.Price, 
		Weight: bookDTO.Detail.Weight,
	}

	if err := models.DB.Save(&savedBook).Error; err != nil {
		return savedBook, err
	}

	return savedBook, nil
}

func (c *BookService) DeleteBook(id string) error {
	var book models.Book;

	if err := models.DB.First(&book, id).Error; err != nil {
		return err
	}

	if err := models.DB.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}