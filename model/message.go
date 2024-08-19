package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Body     string
	BodyType int
	FromId   int
	ToId     int
}

func (Message) TableName() string {
	return "message"
}

func NewMessage(content string, contentType int) *Message {
	return &Message{
		Body:     content,
		BodyType: contentType,
		FromId:   99,
		ToId:     100,
	}
}
