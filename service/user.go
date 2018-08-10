package service

import (
	"monitorGo/model"
)

func (s *Service) UserList() (users []*model.User, err error) {
	users, err = s.dao.UserList()
	return
}

func (s *Service) SaveUser(user *model.User) (err error) {
	err = s.dao.SaveUser(user)
	return
}
