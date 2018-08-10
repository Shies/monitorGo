package service

import (
	"monitorGo/conf"
	"monitorGo/dao"
	"monitorGo/task"
	"time"
	"monitorGo/model"
	"sync"
	"log"
	"strings"
	)

// Service biz service def.
type Service struct {
	c    *conf.Config
	dao	 *dao.Dao
	Test chan string
	Send chan []string
	Quit chan bool
	Sync *sync.WaitGroup
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:    c,
		dao:  dao.New(c),
		Test: make(chan string),
		Send: make(chan []string, 10),
		Quit: make(chan bool, 1),
		Sync: new(sync.WaitGroup),
	}

	go s.loadTask();
	return s
}

func (s *Service) loadTask() {
	for {
		tasks := s.dao.TaskList(dao.TASK_BY_ALL, "1")
		ips := s.dao.TaskIP(dao.IPS_BY_ALL, 1)
		for _, v := range tasks {
			s.Req(v, ips[v.Id])
		}
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func (s *Service) Req(t *model.TaskItem, ips []*model.TaskIP) {
	if ips != nil {
		s.Sync.Add(1)
		go s.Consumer(t)

		go s.Producer(ips)
		go s.Tester()

		go s.Done()
	} else {

		go func() {
			if _, err := task.HttpDo(t.Method, t.Url, t.Params, nil); err != nil {
				log.Printf("%v\n", err)
				return
			}
		}()
	}
	return
}

func (s *Service) Tester() {
	s.Test <- "hello world"
}

func (s *Service) Producer(ips []*model.TaskIP) {
	var ipstr = []string{}
	for _, v := range ips {
		ipstr = append(ipstr, v.IP)
	}
	select {
	case s.Send <- ipstr:
	default:
		log.Println("the chan is full(" + strings.Join(ipstr, ",") + ")")
	}
	return
}

func (s *Service) Consumer(t *model.TaskItem) {
	defer s.Sync.Done()
	for {
		select {
		case ipstr, ok := <-s.Send:
			if !ok {
				return
			}
			var header = make(map[string]string)
			for _, ip := range ipstr {
				log.Println("start:" + t.Url)
				urlData := task.ParseUrl(t.Url, ip)
				part := strings.Split(urlData["header"], ":")
				header["host"] = part[1]
				if _, err := task.HttpDo(t.Method, urlData["url"], t.Params, header); err != nil {
					log.Printf("%v\n", urlData)
					continue
				}
			}
		case welcome := <-s.Test:
			log.Println(welcome)
		case <-s.Quit:
			log.Println("done")
			return
		}
	}
}

func (s *Service) Done() {
	time.Sleep(time.Duration(3) * time.Second)
	s.Quit <- true
}

func (s *Service) Close() {
	close(s.Send)
	close(s.Test)
	close(s.Quit)
}

// Ping check server ok.
func (s *Service) Ping() (err error) {
	return
}