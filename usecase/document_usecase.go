package usecase

import (
	"context"
	"time"

	"github.com/connorb645/appeak-go/domain"
)

type documentUsecase struct {
	documentStore  domain.DocumentStore
	contextTimeout time.Duration
}

func NewDocumentUsecase(documentStore domain.DocumentStore, timeout time.Duration) domain.DocumentUsecase {
	return &documentUsecase{
		documentStore:  documentStore,
		contextTimeout: timeout,
	}
}

func (du *documentUsecase) FetchAllDocuments(c context.Context) ([]domain.Document, error) {
	_, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.documentStore.FetchAll()
}
