package main

import "monitorGo/controller"

func main() {
	/*
		tasks := model.New().TaskList(model.TASK_BY_ALL, "1")
		ips := model.New().TaskIP(model.IPS_BY_ALL, 1)
		for _, v := range tasks {
			task.Request(v, ips[v.Id])
		}
	*/
	controller.Register()
}
