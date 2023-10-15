package server

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

    if err != nil{
        return nil, err
    }

    db.AutoMigrate(&Folder{})
    return db, err
}


type Folder struct {
    gorm.Model
    Id string `gorm:"primarykey"`
    Name string
}


