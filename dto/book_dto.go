package dto

type BookDetailDTO struct {
	Price uint32 `json:"price" validate:"required,min=0"`
	Weight uint32 `json:"weight" validate:"required,min=1"`
}

type BookDTO struct {
	Title string `json:"title" validate:"required,min=5,max=255"`
	Author string `json:"author" validate:"required,alpha,min=5,max=100"`
	Description string `json:"description" validate:"required"`
	PublishDate string `json:"publish_date" validate:"required"`
	Tags []string `json:"tags" validate:"required,gt=0,dive,required"`
	Detail *BookDetailDTO `json:"detail" validate:"required"`
}