package service

import (
	"monitorGo/model"
)

func (s *Service) TaskList(sql string, param string) (tasks []*model.TaskItem) {
	tasks = s.dao.TaskList(sql, param)
	return
}

func (s *Service) SaveTask(taskItem *model.TaskItem) {
	s.dao.SaveTask(taskItem)
	return
}
