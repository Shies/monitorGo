package service

import (
	"monitorGo/model"
	"strings"
	"log"
)

func (s *Service) TaskList(sql string, param string) (tasks []*model.TaskItem, err error) {
	tasks, err = s.dao.TaskList(sql, param)
	return
}

func (s *Service) SaveTask(taskItem *model.TaskItem) (err error) {
	err = s.dao.SaveTask(taskItem)
	return
}

func (s *Service) SendMails(emails []string, msg string) error {
	var (
		err error
		confs, _ = s.dao.ConfList()
	)
	globalEmails := strings.Split(confs["GlobalEmails"], ",")
	for _, email := range globalEmails {
		emails = append(emails, email)
	}

	log.Printf("%v\n", emails)
	for _, email := range emails {
		err = s.mail.Send(email, "告警邮件", msg, "normal")
		if err != nil {
			log.Printf("%v\n", err)
			return err
		}
	}

	return nil
}
