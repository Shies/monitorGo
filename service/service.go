package service

import (
	"monitorGo/conf"
	"monitorGo/model"
)

// Service biz service def.
type Service struct {
	c    *conf.Config
	dao	 *model.Dao
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:         c,
		dao:       model.New(),
	}
	return s
}