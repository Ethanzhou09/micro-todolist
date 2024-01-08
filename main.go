package main

import (
	"log"
	api "todolist/idl/task/kitex_gen/api/taskservice"
)

func main() {
	svr := api.NewServer(new(TaskServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
