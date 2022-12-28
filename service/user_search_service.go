package service

import (
	"errors"
	"log"
	"strconv"
	"time"

	"go-redisearch/domain"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type UserSearchService struct {
	redisearchclient *redisearch.Client
}

type SearchService interface {
	InsertUserDocument(domain.User) string
	GetUserDocumentByFirstname(firstname string) (redisearch.Document, error)
	DeleteUserDocument(docId string) bool
}

func NewUserSearchService(redisearchclient *redisearch.Client) SearchService {
	return &UserSearchService{redisearchclient: redisearchclient}
}

func (u *UserSearchService) InsertUserDocument(user domain.User) string {
	unixTime := time.Now().UTC().Unix()
	id := strconv.Itoa(int(unixTime))

	doc := redisearch.NewDocument(id, 1.0)
	doc.Set("firstname", user.Firstname).
		Set("lastname", user.Lastname).
		Set("date", time.Now().Unix())

	if err := u.redisearchclient.Index([]redisearch.Document{doc}...); err != nil {
		log.Fatal(err)
	}

	return id
}

func (u *UserSearchService) GetUserDocumentByFirstname(firstname string) (redisearch.Document, error) {
	docs, total, err := u.redisearchclient.Search(redisearch.NewQuery(firstname).
		Limit(0, 2).
		SetReturnFields("firstname"))

	//log.Println(docs[0].Id, docs[0].Properties["firstname"], total, err)
	if total > 0 && err == nil {
		return docs[0], nil
	}
	return redisearch.Document{}, errors.New("No document")
}

func (u *UserSearchService) DeleteUserDocument(docId string) bool {
	err := u.redisearchclient.Delete(docId, true)
	if err != nil {
		log.Println("delete error")
		return false
	}
	return true
}
