// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/connorb645/appeak-go/domain"
	mock "github.com/stretchr/testify/mock"
)

// DocumentUsecase is an autogenerated mock type for the DocumentUsecase type
type DocumentUsecase struct {
	mock.Mock
}

// FetchAllDocuments provides a mock function with given fields: c
func (_m *DocumentUsecase) FetchAllDocuments(c context.Context) ([]domain.Document, error) {
	ret := _m.Called(c)

	var r0 []domain.Document
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Document, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Document); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Document)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDocumentUsecase creates a new instance of DocumentUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDocumentUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *DocumentUsecase {
	mock := &DocumentUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
