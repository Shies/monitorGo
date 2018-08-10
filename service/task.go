package service

import (
	"monitorGo/model"
)

func (s *Service) TaskList(sql string, param string) (tasks []*model.TaskItem, err error) {
	tasks, err = s.dao.TaskList(sql, param)
	return
}

func (s *Service) SaveTask(taskItem *model.TaskItem) (err error) {
	err = s.dao.SaveTask(taskItem)
	return
}
