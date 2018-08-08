package service

import (
	"monitorGo/model"
)

func (s *Service) UserList() (users []*model.User) {
	users = s.dao.UserList()
	return
}

func (s *Service) SaveUser(user *model.User) {
	s.dao.SaveUser(user)
	return
}
