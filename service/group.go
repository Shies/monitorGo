package service

import (
	"monitorGo/model"
)

func (s *Service) GroupList() (group []*model.Group, err error) {
	group, err = s.dao.GroupList()
	return
}

func (s *Service) SaveGroup(group *model.Group) (err error) {
	err = s.dao.SaveGroup(group)
	return
}