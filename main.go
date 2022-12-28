package main

import (
	"go-redisearch/config"
	"go-redisearch/domain"
	"go-redisearch/service"
	"log"
)

func main() {

	redisearchConfig := config.RedisearchConfig()
	userSearchService := service.NewUserSearchService(redisearchConfig)

	user := domain.User{Firstname: "mert", Lastname: "çakmak"}

	userId := userSearchService.InsertUserDocument(user)
	log.Printf("User document inserted. ID: %s", userId)

	userDoc, err := userSearchService.GetUserDocumentByFirstname(user.Firstname)
	if err == nil {
		log.Printf("Retrieved user by firstname. ID: %s, firstname: %s", userDoc.Id, userDoc.Properties["firstname"])
	}

	isDeleted := userSearchService.DeleteUserDocument(userDoc.Id)
	if isDeleted {
		log.Printf("Deleted user by id. ID: %s", userDoc.Id)
	}
}
