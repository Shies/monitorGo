package service

import (
	"monitorGo/model"
)

func (s *Service) TaskIP(query string, param int64) (ips map[int64][]*model.TaskIP) {
	ips = s.dao.TaskIP(query, param)
	return
}

func (s *Service) SaveIP(ip *model.TaskIP) {
	s.dao.SaveIP(ip)
	return
}
