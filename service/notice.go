package service

import (
	"monitorGo/model"
)

func (s *Service) SendList(sql string, param int64) (send map[int64][]*model.Notice) {
	send = s.dao.SendList(sql, param)
	return
}

func (s *Service) SaveNotice(notice *model.Notice) {
	s.dao.SaveNotice(notice)
	return
}
