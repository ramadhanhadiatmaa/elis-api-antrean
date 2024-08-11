package models

import "gorm.io/gorm"

type Ant struct {
	gorm.Model
	Seq int `json:"ant"`
}
