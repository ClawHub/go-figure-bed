package models

//公共model
type Model struct {
	ID         string `gorm:"primary_key" json:"id"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`
}
