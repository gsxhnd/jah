package model

import (
	"time"
)

type Anime struct {
	Id          uint       `json:"id"`
	Code        string     `json:"code"`
	Title       string     `json:"title"`
	TitleCn     string     `json:"title_cn"`
	Cover       *string    `json:"cover"`
	PublishDate *time.Time `json:"publish_date"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
