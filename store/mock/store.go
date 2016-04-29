package mock

import "github.com/stretchr/testify/mock"

import "github.com/lgtmco/lgtm/model"

type Store struct {
	mock.Mock
}

func (_m *Store) GetUser(_a0 int64) (*model.User, error) {
	ret := _m.Called(_a0)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(int64) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Store) GetUserLogin(_a0 string) (*model.User, error) {
	ret := _m.Called(_a0)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Store) CreateUser(_a0 *model.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Store) UpdateUser(_a0 *model.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Store) DeleteUser(_a0 *model.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Store) GetRepo(_a0 int64) (*model.Repo, error) {
	ret := _m.Called(_a0)

	var r0 *model.Repo
	if rf, ok := ret.Get(0).(func(int64) *model.Repo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Store) GetRepoSlug(_a0 string) (*model.Repo, error) {
	ret := _m.Called(_a0)

	var r0 *model.Repo
	if rf, ok := ret.Get(0).(func(string) *model.Repo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Store) GetRepoMulti(_a0 ...string) ([]*model.Repo, error) {
	ret := _m.Called(_a0)

	var r0 []*model.Repo
	if rf, ok := ret.Get(0).(func(...string) []*model.Repo); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Store) GetRepoOwner(_a0 string) ([]*model.Repo, error) {
	ret := _m.Called(_a0)

	var r0 []*model.Repo
	if rf, ok := ret.Get(0).(func(string) []*model.Repo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Store) CreateRepo(_a0 *model.Repo) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Repo) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Store) UpdateRepo(_a0 *model.Repo) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Repo) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Store) DeleteRepo(_a0 *model.Repo) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Repo) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}