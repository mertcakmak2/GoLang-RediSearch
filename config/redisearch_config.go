package config

import (
	"log"

	"github.com/RediSearch/redisearch-go/redisearch"
)

func RedisearchConfig() *redisearch.Client {
	redisearchClient := redisearch.NewClient("localhost:6379", "userIndex")

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("firstname")).
		AddField(redisearch.NewTextFieldOptions("lastname", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("date"))

	redisearchClient.Drop()

	if err := redisearchClient.CreateIndex(sc); err != nil {
		log.Fatal(err)
	}

	return redisearchClient
}
