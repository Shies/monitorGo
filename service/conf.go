package service

func (s *Service) ConfList() (map[string]string) {
	return s.dao.ConfList()
}

func (s *Service) SaveConf(key string, val string) {
	s.dao.SaveConf(key, val)
	return
}
