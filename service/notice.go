package service

import (
	"monitorGo/model"
)

func (s *Service) SendList(sql string, param int64) (send map[int64][]*model.Notice, err error) {
	send, err = s.dao.SendList(sql, param)
	return
}

func (s *Service) SaveNotice(notice *model.Notice) (err error) {
	err = s.dao.SaveNotice(notice)
	return
}
