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
type Service struct {
	c    *conf.Config
	dao	 *dao.Dao
	wait *sync.WaitGroup
	once sync.Once
	test chan string
	quit chan bool
	send map[int64]chan []*model.TaskIP
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:    c,
		dao:  dao.New(c),
		wait: new(sync.WaitGroup),
		test: make(chan string),
		quit: make(chan bool, 1),
		send: make(map[int64]chan []*model.TaskIP),
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
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func (s *Service) R(tis []*model.TaskItem, ips map[int64][]*model.TaskIP) {
	var (
		tmp = make(map[int64]*model.TaskItem)
		diff []*model.TaskItem
	)
	for _, t := range tis {
		if _, ok := ips[t.Id]; ok {
			tmp[t.Id] = t
		} else {
			diff = append(diff, t)
		}
	}
	if tmp != nil {
		s.wait.Add(2)
		go s.Consumer(tmp)
		go s.Producer(ips)
		// s.Wait.Wait()

		s.once.Do(func() {
			tmp = nil
			ips = nil
		})
	}

	go func() {
		for _, t := range diff {
			if _, err := task.HttpDo(t.Method, t.Url, t.Params, nil); err != nil {
				log.Printf("%v\n", err)
				return
			}
		}
	}()
	return
}

func (s *Service) Tester() {
	s.test <- "hello world"
}

func (s *Service) Producer(ips map[int64][]*model.TaskIP) {
	defer s.wait.Done()
	for tid, v := range ips {
		var tmp = make(chan []*model.TaskIP, len(v))
		select {
		case tmp <- v:
			s.send[tid] = tmp
		default:
			for _, ip := range v {
				log.Printf("%s", "the chan is full(" + ip.IP + ")")
			}
		}
	}
	go s.Tester()
	return
}

func (s *Service) Consumer(tis map[int64]*model.TaskItem) {
	defer s.wait.Done()
	for {
		for tid, v := range s.send {
			select {
			case ips, ok := <-v:
				if !ok {
					s.Close()
					return
				}

				t := tis[tid]
				delete(s.send, tid)
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
			}
		}
		select {
		case welcome := <-s.test:
			log.Println(welcome)
		case <-time.After(time.Duration(3) * time.Microsecond):
			s.Done()
		case <-s.quit:
			log.Println("done")
			return
		}
	}
}

func (s *Service) Done() {
	s.quit <- true
}

func (s *Service) Close() {
	for _, v := range s.send {
		close(v)
	}
	close(s.test)
	close(s.quit)
}

// Ping check server ok.
func (s *Service) Ping() (err error) {
	return
}