package domain

import "context"

type Document struct {
	ID       int 	`json:"id"`
	Title    string `json:"title"`
	BodyHTML string `json:"body"`
}

type DocumentStore interface {
	FetchAll() ([]Document, error)
	Fetch(id string) (*Document, error)
}

type DocumentUsecase interface {
	FetchAllDocuments(c context.Context) ([]Document, error)
}
