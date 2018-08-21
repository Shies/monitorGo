package service

import (
	"time"
	"sync"
	"log"
	"strings"

	"monitorGo/conf"
	"monitorGo/dao"
	"monitorGo/task"
	"monitorGo/model"
)


// Service biz service def.
type Service2 struct {
	c    *conf.Config
	dao	 *dao.Dao
	wait *sync.WaitGroup
	once sync.Once
	test chan string
	quit chan bool
	send chan []*model.TaskIP
}

// New new a Service and return.
func New2(c *conf.Config) (s *Service2) {
	s = &Service2{
		c:    c,
		dao:  dao.New(c),
		wait: new(sync.WaitGroup),
		test: make(chan string),
		quit: make(chan bool, 1),
		send: make(chan []*model.TaskIP, 10240),
	}

	go s.loadTaskTick();
	return s
}

func (s *Service2) loadTaskTick() {
	for {
		tasks, _ := s.dao.TaskList(dao.TASK_BY_ALL, "1")
		ips, _ := s.dao.TaskIP(dao.IPS_BY_ALL, 1)
		for _, t := range tasks {
			s.R(t, ips[t.Id])
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func (s *Service2) R(t *model.TaskItem, ips []*model.TaskIP) {
	if ips != nil {
		s.wait.Add(2)
		go s.Consumer(t)
		go s.Producer(ips)
		// s.wait.Wait()

		s.once.Do(func() {
			t = nil
			ips = nil
		})
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

func (s *Service2) Tester() {
	s.test <- "hello world"
}

func (s *Service2) Producer(ips []*model.TaskIP) {
	defer s.wait.Done()
	select {
	case s.send <- ips:
	default:
		for _, ip := range ips {
			log.Printf("%s", "the chan is full(" + ip.IP + ")")
		}
	}
	go s.Tester()
	return
}

func (s *Service2) Consumer(t *model.TaskItem) {
	defer s.wait.Done()
	for {
		select {
		case ips, ok := <-s.send:
			if !ok {
				s.Close()
				return
			}

			var header = make(map[string]string)
			for _, ip := range ips {
				log.Println("start:" + t.Url)
				urlData := task.ParseUrl(t.Url, ip.IP)
				part := strings.Split(urlData["header"], ":")
				header["host"] = part[1]
				if _, err := task.HttpDo(t.Method, urlData["url"], t.Params, header); err != nil {
					log.Printf("%v\n", urlData)
					continue
				}
			}
		case welcome := <-s.test:
			log.Println(welcome)
		case <-time.After(time.Duration(1) * time.Second):
			s.Done()
		case <-s.quit:
			log.Println("done")
			return
		}
	}
}

func (s *Service2) Done() {
	s.quit <- true
}

func (s *Service2) Close() {
	close(s.send)
	close(s.test)
	close(s.quit)
}

// Ping check server ok.
func (s *Service2) Ping() (err error) {
	return
}