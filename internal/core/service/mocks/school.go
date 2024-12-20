// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/paw1a/eschool/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// SchoolRepository is an autogenerated mock type for the ISchoolRepository type
type SchoolRepository struct {
	mock.Mock
}

// AddSchoolTeacher provides a mock function with given fields: ctx, schoolID, teacherID
func (_m *SchoolRepository) AddSchoolTeacher(ctx context.Context, schoolID domain.ID, teacherID domain.ID) error {
	ret := _m.Called(ctx, schoolID, teacherID)

	if len(ret) == 0 {
		panic("no return value specified for AddSchoolTeacher")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID, domain.ID) error); ok {
		r0 = rf(ctx, schoolID, teacherID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, school
func (_m *SchoolRepository) Create(ctx context.Context, school domain.School) (domain.School, error) {
	ret := _m.Called(ctx, school)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 domain.School
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.School) (domain.School, error)); ok {
		return rf(ctx, school)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.School) domain.School); ok {
		r0 = rf(ctx, school)
	} else {
		r0 = ret.Get(0).(domain.School)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.School) error); ok {
		r1 = rf(ctx, school)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, schoolID
func (_m *SchoolRepository) Delete(ctx context.Context, schoolID domain.ID) error {
	ret := _m.Called(ctx, schoolID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) error); ok {
		r0 = rf(ctx, schoolID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx
func (_m *SchoolRepository) FindAll(ctx context.Context) ([]domain.School, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []domain.School
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.School, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.School); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.School)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, schoolID
func (_m *SchoolRepository) FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error) {
	ret := _m.Called(ctx, schoolID)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 domain.School
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) (domain.School, error)); ok {
		return rf(ctx, schoolID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) domain.School); ok {
		r0 = rf(ctx, schoolID)
	} else {
		r0 = ret.Get(0).(domain.School)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, schoolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSchoolCourses provides a mock function with given fields: ctx, schoolID
func (_m *SchoolRepository) FindSchoolCourses(ctx context.Context, schoolID domain.ID) ([]domain.Course, error) {
	ret := _m.Called(ctx, schoolID)

	if len(ret) == 0 {
		panic("no return value specified for FindSchoolCourses")
	}

	var r0 []domain.Course
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) ([]domain.Course, error)); ok {
		return rf(ctx, schoolID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) []domain.Course); ok {
		r0 = rf(ctx, schoolID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Course)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, schoolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSchoolTeachers provides a mock function with given fields: ctx, schoolID
func (_m *SchoolRepository) FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error) {
	ret := _m.Called(ctx, schoolID)

	if len(ret) == 0 {
		panic("no return value specified for FindSchoolTeachers")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) ([]domain.User, error)); ok {
		return rf(ctx, schoolID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) []domain.User); ok {
		r0 = rf(ctx, schoolID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, schoolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserSchools provides a mock function with given fields: ctx, userID
func (_m *SchoolRepository) FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for FindUserSchools")
	}

	var r0 []domain.School
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) ([]domain.School, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) []domain.School); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.School)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsSchoolTeacher provides a mock function with given fields: ctx, schoolID, teacherID
func (_m *SchoolRepository) IsSchoolTeacher(ctx context.Context, schoolID domain.ID, teacherID domain.ID) (bool, error) {
	ret := _m.Called(ctx, schoolID, teacherID)

	if len(ret) == 0 {
		panic("no return value specified for IsSchoolTeacher")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID, domain.ID) (bool, error)); ok {
		return rf(ctx, schoolID, teacherID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID, domain.ID) bool); ok {
		r0 = rf(ctx, schoolID, teacherID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID, domain.ID) error); ok {
		r1 = rf(ctx, schoolID, teacherID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, school
func (_m *SchoolRepository) Update(ctx context.Context, school domain.School) (domain.School, error) {
	ret := _m.Called(ctx, school)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 domain.School
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.School) (domain.School, error)); ok {
		return rf(ctx, school)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.School) domain.School); ok {
		r0 = rf(ctx, school)
	} else {
		r0 = ret.Get(0).(domain.School)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.School) error); ok {
		r1 = rf(ctx, school)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSchoolRepository creates a new instance of SchoolRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSchoolRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SchoolRepository {
	mock := &SchoolRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
