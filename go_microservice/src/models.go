package main

import (
	"github.com/google/uuid"
	"time"
)

type Greet struct {
	Id        uuid.UUID `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	Greet     string    `gorm:"column:greet"`
}

//type jsonTime time.Time
