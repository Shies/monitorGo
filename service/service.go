package service

import (
	"monitorGo/conf"
	"monitorGo/dao"
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
	return s
}


func (s *Service) Close() {
	return
}

// Ping check server ok.
func (s *Service) Ping() (err error) {
	return
}