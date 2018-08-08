package service

import (
	"monitorGo/model"
)

func (s *Service) FaultList(tid int64, ip string) (faults []*model.Fault) {
	faults = s.dao.FaultList(tid, ip)
	return
}

func (s *Service) FaultTid() (tid int64) {
	tid = s.dao.FaultTid()
	return
}