package store

import (
	"github.com/connorb645/appeak-go/domain"
)

type documentStore struct {
	helpCenterProvider HelpCenterProvider
}

type HelpCenterProvider interface {
	GetDocuments() ([]domain.Document, error)
	GetDocument(id string) (*domain.Document, error)
}

func NewDocumentStore(hcp HelpCenterProvider) domain.DocumentStore {
	return &documentStore{
		helpCenterProvider: hcp,
	}
}

func (dr *documentStore) FetchAll() ([]domain.Document, error) {
	documents, err := dr.helpCenterProvider.GetDocuments()
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (dr *documentStore) Fetch(id string) (*domain.Document, error) {
	doc, err := dr.helpCenterProvider.GetDocument(id)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
