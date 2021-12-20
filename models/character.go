package models

type Character struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"title"`
	Class string `json:"class"`
	Level uint8  `json:"level"`
}

type CreateCharacterInput struct {
	Name  string `json:"name" binding:"required"`
	Class string `json:"class" binding:"required"`
}

type UpdateCharacterInput struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}
