package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NotesAll() *[]Note {
	var notes []Note
	DB.Where("deleted_at is NULL").Order("updated_at desc").Find(&notes)
	return &notes
}

func NoteCreate(name string, content string) *Note {
	entry := Note{Name: name, Content: content}
	DB.Create(&entry)
	return &entry
}

func NotesFind(id uint64) *Note {
	var note Note
	DB.Where("id =?", id).First(&note)
	return &note
}

func (note *Note) Update(name string, content string) {
	note.Name = name
	note.Content = content
	DB.Save(note)
}
