package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LessonRecord struct {
	ObjID     primitive.ObjectID `bson:"_id,omitempty"`
	User      User               `bson:"user"`
	Lesson    Lesson             `bson:"lesson"`
	Other     Other              `bson:"other"`
	Type      int                `bson:"type"`
	Sample    bool               `bson:"sample"`
	Result    int                `bson:"result"`
	Status    bool               `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
	V         int                `bson:"__v"`
}
type Other struct {
	ID primitive.ObjectID `bson:"_id"`
}
type Lesson struct {
	Parent    Parent             `bson:"parent"`
	GG        int                `bson:"gg"`
	Level     int                `bson:"level"`
	Tittle    string             `bson:"tittle"`
	Thumbnail string             `bson:"thumbnail"`
	ID        primitive.ObjectID `bson:"_id"`
	Sentences []Sentence         `bson:"sentences"`
}
type Sentence struct {
	User     UserSentence       `bson:"user"`
	Record   Record             `bson:"record"`
	ID       primitive.ObjectID `bson:"_id"`
	Sentence string             `bson:"sentence"`
	Subtitle string             `bson:"subtitle"`
	Duration float64            `bson:"duration"`
	Audio    string             `bson:"audio"`
}
type Record struct {
	GG     int    `bson:"gg"`
	Result Result `bson:"result"`
	Audio  string `bson:"audio"`
}
type Result struct {
	NumDeletions     int         `bson:"numDeletions"`
	NumInsertions    int         `bson:"numInsertions"`
	NumSubstitutions int         `bson:"numSubstitutions"`
	TextHypothesis   string      `bson:"textHypothesis"`
	TextReference    string      `bson:"textReference"`
	TotalErrors      int         `bson:"totalErrors"`
	TotalWords       int         `bson:"totalWords"`
	Wer              interface{} `bson:"wer"`
	Words            []Word      `bson:"words"`
}
type Word struct {
	ID      primitive.ObjectID `bson:"_id"`
	Index   int                `bson:"index"`
	Type    string             `bson:"string"`
	WordHyp string             `bson:"wordHyp"`
	WordRef string             `bson:"wordRef"`
}
type UserSentence struct {
	Name   string `bson:"name"`
	Avatar string `bson:"avatar"`
	Gender bool   `bson:"gender"`
}
type Parent struct {
	Title     string `bson:"title"`
	Slug      string `bson:"slug"`
	Thumbnail string `bson:"thumbnail"`
}
type User struct {
	UserID           int     `bson:"user_id"`
	GG               int     `bson:"gg"`
	Accuracy         float64 `bson:"accuracy"`
	CurrrentSentence string  `bson:"current_sentence"`
}
