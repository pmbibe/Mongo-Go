package main

import (
	"log"
)

func main() {
	pCollection, ctx := connectMongo(db, collection, url)
	cursor, err := pCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &recordAudios); err != nil {
		log.Fatal(err)
	}
	for _, recordAudio := range recordAudios {
		for _, sentence := range recordAudio.Lesson.Sentences {
			if re.MatchString(sentence.Record.Audio) {
				updateMongo(pCollection, ctx, newURL, oldURL, sentence, key)
			}
		}
	}
}
