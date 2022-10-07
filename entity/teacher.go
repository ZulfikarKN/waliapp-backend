package entity

import "time"

type Teacher struct {
	ID              int       `gorm:"column:id;primary_key" json:"id"`
	Name            string    `gorm:"column:name" json:"name"`
	Image           string    `gorm:"column:image" json:"image"`
	Studies         string    `gorm:"column:studies" json:"studies"`
	HomeroomTeacher string    `gorm:"column:homeroom_teacher" json:"homeroom_teacher"`
	Join_date       time.Time `gorm:"column:join_date" json:"join_date"`
	Status 			int 	  `gorm:"column:status" json:"status"`
}

type TeacherViewModel struct {
	ID              int      `gorm:"column:id;primary_key" json:"id"`
	Name            string   `gorm:"column:name" json:"name"`
	Image           string   `gorm:"column:image" json:"image"`
	Studies         string `gorm:"column:studies" json:"studies"`
	HomeroomTeacher string   `gorm:"column:homeroom_teacher" json:"homeroom_teacher"`
	Status 			int 	 `gorm:"column:status" json:"status"`
}