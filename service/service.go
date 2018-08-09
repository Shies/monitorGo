package service

import (
	"monitorGo/conf"
	"monitorGo/dao"
	"monitorGo/task"
	"time"
)

// Service biz service def.
type Service struct {
	c    *conf.Config
	dao	 *dao.Dao
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:         c,
		dao:       dao.New(),
	}

	go s.asyncTask();
	return s
}

func (s *Service) asyncTask() {
	for {
		tasks := s.dao.TaskList(dao.TASK_BY_ALL, "1")
		ips := s.dao.TaskIP(dao.IPS_BY_ALL, 1)
		for _, v := range tasks {
			task.Request(v, ips[v.Id])
		}
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func (s *Service) Close() {
	return
}

// Ping check server ok.
func (s *Service) Ping() (err error) {
	return
}