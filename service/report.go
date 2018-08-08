package service

import (
	"monitorGo/model"
)

func (s *Service) ReportList(tid int64, ip string) (reports []*model.Report) {
	reports = s.dao.ReportList(tid, ip)
	return
}

func (s *Service) ReportTid() (tid int64) {
	tid = s.dao.ReportTid()
	return
}

func (s *Service) StateReport() (state []*model.Status) {
	state = s.dao.StateReport()
	return
}

func (s *Service) IndexReport() (index []*model.Index) {
	index = s.dao.IndexReport()
	return
}