package service

import (
	"monitorGo/conf"
	"monitorGo/dao"
	"monitorGo/task"
	"monitorGo/model"

	"time"
	"sync"
	"log"
	"strings"
)

const (
	_sharding = 10240
)

// Service biz service def.
type Service struct {
	c    *conf.Config
	dao	 *dao.Dao
	Wait *sync.WaitGroup
	Once sync.Once
	Test chan string
	Quit chan bool
	Send chan map[int64][]*model.TaskIP
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:    c,
		dao:  dao.New(c),
		Wait: new(sync.WaitGroup),
		Test: make(chan string),
		Quit: make(chan bool, 1),
		Send: make(chan map[int64][]*model.TaskIP, _sharding),
	}

	go s.loadTaskTick();
	return s
}

func (s *Service) loadTaskTick() {
	for {
		tasks, _ := s.dao.TaskList(dao.TASK_BY_ALL, "1")
		ips, _ := s.dao.TaskIP(dao.IPS_BY_ALL, 1)
		if tasks != nil || ips != nil {
			s.R(tasks, ips)
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func (s *Service) R(tis []*model.TaskItem, ips map[int64][]*model.TaskIP) {
	if ips != nil {
		s.Wait.Add(2)
		go s.Consumer(tis)
		go s.Producer(ips)
		s.Wait.Wait()

		s.Once.Do(func() {
			tis = nil
			ips = nil
		})
	} else {
		go func() {
			for _, t := range tis {
				if _, err := task.HttpDo(t.Method, t.Url, t.Params, nil); err != nil {
					log.Printf("%v\n", err)
					return
				}
			}
		}()
	}
	return
}

func (s *Service) Tester() {
	s.Test <- "hello world"
}

func (s *Service) Producer(ips map[int64][]*model.TaskIP) {
	defer s.Wait.Done()

	select {
	case s.Send <- ips:
	default:
		log.Printf("%s%v", "the chan is full", ips)
	}

	go s.Tester()
	go s.Done()
	return
}

func (s *Service) Consumer(tis []*model.TaskItem) {
	defer s.Wait.Done()
	for {
		select {
		case ips, ok := <-s.Send:
			if !ok {
				return
			}
			for _, t := range tis {
				var header = make(map[string]string)
				for _, ip := range ips[t.Id] {
					log.Println("start:" + t.Url)
					urlData := task.ParseUrl(t.Url, ip.IP)
					part := strings.Split(urlData["header"], ":")
					header["host"] = part[1]
					if _, err := task.HttpDo(t.Method, urlData["url"], t.Params, header); err != nil {
						log.Printf("%v\n", urlData)
						continue
					}
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