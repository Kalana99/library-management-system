package models

import "gorm.io/gorm"

type Issue struct {
    gorm.Model
    UserID uint `json:"user_id"`
    BookID uint `json:"book_id"`
    Returned bool `json:"returned"`
}
