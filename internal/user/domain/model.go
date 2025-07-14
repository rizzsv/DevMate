package domain 

import "time"

type User struct {
	ID string `gorm:"primaryKey"`
	Name string 
	Email string `gorm:"unique"`
	Password string
	Role string
	Level string `gorm: "type:VARCHAR(20);check:level IN ('basic', 'intermediate', 'advanced')"`
	Bio *string
	CreatedAt time.Time
	UpdatedAt time.Time
}