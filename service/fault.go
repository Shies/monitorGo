package service

import (
	"monitorGo/model"
)

func (s *Service) FaultList(tid int64, ip string) (faults []*model.Fault, err error) {
	faults, err = s.dao.FaultList(tid, ip)
	return
}

func (s *Service) FaultTid() (tid int64, err error) {
	tid, err = s.dao.FaultTid()
	return
}