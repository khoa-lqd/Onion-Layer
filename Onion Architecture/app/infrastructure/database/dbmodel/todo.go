package dbmodel

import "time"

type Todo struct {
	ID        int       `gorm:"primary_key;not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Complete  bool      `gorm:"type:boolean;not null"`
	Deadline  time.Time `gorm:"type:datetime;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

func (t Todo) TableName() string {
	return "todo"
}
