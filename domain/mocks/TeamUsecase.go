// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/connorb645/appeak-go/domain"
	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// TeamUsecase is an autogenerated mock type for the TeamUsecase type
type TeamUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, team
func (_m *TeamUsecase) Create(c context.Context, team *domain.Team) error {
	ret := _m.Called(c, team)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Team) error); ok {
		r0 = rf(c, team)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: c, teamID
func (_m *TeamUsecase) Fetch(c context.Context, teamID primitive.ObjectID) (*domain.Team, error) {
	ret := _m.Called(c, teamID)

	var r0 *domain.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (*domain.Team, error)); ok {
		return rf(c, teamID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Team); ok {
		r0 = rf(c, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(c, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, teamID, team
func (_m *TeamUsecase) Update(c context.Context, teamID primitive.ObjectID, team *domain.TeamUpdate) error {
	ret := _m.Called(c, teamID, team)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.TeamUpdate) error); ok {
		r0 = rf(c, teamID, team)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTeamUsecase creates a new instance of TeamUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTeamUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *TeamUsecase {
	mock := &TeamUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
