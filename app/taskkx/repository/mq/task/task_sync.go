package task

import (
	"context"
	"encoding/json"
	"todolist/app/taskkx/repository/mq"
	"todolist/app/taskkx/service"
	"todolist/consts"
	"todolist/idl/task/kitex_gen/api"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskService(ctx context.Context) (err error) {
	rabbitMqQueue := consts.RabbitMqTaskQueue
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		return
	}
	//var forever chan struct{}
	go func() {
		for d := range msgs {
			// 1.解析消息
			req := new(api.TaskRequest)
			err = json.Unmarshal(d.Body, req)
			if err != nil {
				return
			}
			// 2.落库
			err = service.TaskMq2DB(ctx, req)
			d.Ack(false)
		}
	}()
	// 5.阻塞
	//<-forever
	return nil
}
