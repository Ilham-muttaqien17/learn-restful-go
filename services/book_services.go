package services

import (
	"math"

	"github.com/Ilham-muttaqien17/learn-restful-go/config"
	"github.com/Ilham-muttaqien17/learn-restful-go/dto"
	"github.com/Ilham-muttaqien17/learn-restful-go/models"
	"github.com/Ilham-muttaqien17/learn-restful-go/utils"
	"gorm.io/gorm/clause"
)

type BookService struct{}

func (c *BookService) GetAllBooks(query *utils.PaginationParams) utils.PaginationResponse[[]models.Book] {
	var books []models.Book
	var totalBooks int64

	config.DB.Model(&models.Book{}).Count(&totalBooks)

	config.DB.Offset(int(query.Offset)).Limit(int(query.Limit)).Order(clause.OrderByColumn{
		Column: clause.Column{Name: query.SortBy}, Desc: query.IsDesc,
	}).Find(&books)

	data := utils.ToSlice[models.Book](&books, true)
	// Return empty array/slice instead of null
	if data == nil {
		data = []models.Book{}
	}

	totalPages := int64(math.Ceil(float64(totalBooks) / float64(query.Limit)))

	return utils.PaginationResponse[[]models.Book]{
		Data: data,
		Meta: &utils.MetaPagination{
			Limit:     query.Limit,
			Page:      query.Page,
			TotalData: totalBooks,
			TotalPage: totalPages,
		},
	}
}

func (c *BookService) GetDetailBook(id string) (models.Book, error) {

	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
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
		Price:  bookDTO.Detail.Price,
		Weight: bookDTO.Detail.Weight,
	}

	if err := config.DB.Create(&savedBook).Error; err != nil {
		return savedBook, err
	}

	return savedBook, nil
}

func (c *BookService) UpdateBook(bookDTO dto.BookDTO, id string) (models.Book, error) {
	var savedBook models.Book

	if err := config.DB.First(&savedBook, id).Error; err != nil {
		return savedBook, err
	}

	savedBook.Title = bookDTO.Title
	savedBook.Description = bookDTO.Description
	savedBook.Author = bookDTO.Author
	savedBook.PublishDate = bookDTO.PublishDate
	savedBook.Tags = bookDTO.Tags
	savedBook.Detail = &models.Detail{
		Price:  bookDTO.Detail.Price,
		Weight: bookDTO.Detail.Weight,
	}

	if err := config.DB.Save(&savedBook).Error; err != nil {
		return savedBook, err
	}

	return savedBook, nil
}

func (c *BookService) DeleteBook(id string) error {
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
