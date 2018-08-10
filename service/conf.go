package service

func (s *Service) ConfList() (conf map[string]string, err error) {
	conf, err = s.dao.ConfList()
	return
}

func (s *Service) SaveConf(key string, val string) (err error) {
	err = s.dao.SaveConf(key, val)
	return
}
