package models

type Detail struct {
	Price uint32 `gorm:"uint32" json:"price"`
	Weight uint32 `gorm:"uint32" json:"weight"`
}

type Book struct {
	Id int64 `gorm:"primary" json:"id"`
	Title string `gorm:"type:varchar(300)" json:"title"`
	Author string `gorm:"type:varchar(300)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
	PublishDate string `gorm:"type:date" json:"publish_date"`
	Tags []string `gorm:"serializer:json" json:"tags"`
	Detail *Detail `gorm:"serializer:json" json:"detail"`
}