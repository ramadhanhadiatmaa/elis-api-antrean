package models

import "gorm.io/gorm"

type Antrian struct {
	gorm.Model
	Seq int `json:"ant"`
}
