package service

import (
	"monitorGo/model"
)

func (s *Service) GroupList() (group []*model.Group) {
	group = s.dao.GroupList()
	return
}

func (s *Service) SaveGroup(group *model.Group) {
	s.dao.SaveGroup(group)
	return
}