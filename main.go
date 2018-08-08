package main

import (
	"monitorGo/task"
	"monitorGo/controller"
	"monitorGo/dao"
)

func main() {

	tasks := dao.New().TaskList(dao.TASK_BY_ALL, "1")
	ips := dao.New().TaskIP(dao.IPS_BY_ALL, 1)
	for _, v := range tasks {
		task.Request(v, ips[v.Id])
	}

	controller.Register()
}
