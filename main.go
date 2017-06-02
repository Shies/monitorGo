package main

import "monitorGo/controller"

func main() {
	/*
		tasks := model.GetTask(model.TASK_BY_ALL, "1")
		ips := model.GetTaskIP(model.IPS_BY_ALL, 1)
		for _, v := range tasks {
			task.Request(v, ips[v.Id])
		}
	*/
	controller.Register()
}
