package main

import "log"

func main() {
	pCollection, ctx := connectMongo(db, collection, url)
	cursor, err := pCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &lessons); err != nil {
		log.Fatal(err)
	}
	for _, lesson := range lessons {
		for _, sentence := range lesson.Sentences {
			if re.MatchString(sentence.Audio) {
				updateMongo(pCollection, ctx, newURL, oldURL, sentence, key)
			}
		}
	}
}
