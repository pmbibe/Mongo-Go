package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lesson struct {
	ObjID       primitive.ObjectID `bson:"_id,omitempty"`
	Parent      Parent             `bson:"parent"`
	Title       string             `bson:"title"`
	Thumbnail   string             `bson:"thumbnail"`
	PlayTimes   int                `bson:"play_times"`
	Level       int                `bson:"level"`
	Status      bool               `bson:"status"`
	WinPoint    int                `bson:"win_point"`
	LosePoint   int                `bson:"lose_point"`
	Sentences   []Sentence         `bson:"sentences"`
	CreatedAt   time.Time          `bson:"created_at"`
	V           int                `bson:"__v"`
	GG          int                `bson:"gg"`
	IsPublished bool               `bson:"is_published"`
	IsPaid      bool               `bson:"is_paid"`
}
type Sentence struct {
	User       User               `bson:"user"`
	Sentence   string             `bson:"sentence"`
	Subtitle   string             `bson:"subtitle"`
	Duration   int                `bson:"duration"`
	Audio      string             `bson:"audio"`
	ID         primitive.ObjectID `bson:"_id"`
	SubtitleBR string             `bson:"subtitle_br"`
	SubtitleID string             `bson:"subtitle_id"`
	SubtitleTH string             `bson:"subtitle_th"`
	SubtitleVI string             `bson:"subtitle_vi"`
	SubtitleJA string             `bson:"subtitle_ja"`
}
type User struct {
}
type Parent struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Slug      string             `bson:"slug"`
	Thumbnail string             `bson:"thumbnail"`
	Type      int                `bson:"type"`
	PopupType string             `bson:"popup_type"`
}
