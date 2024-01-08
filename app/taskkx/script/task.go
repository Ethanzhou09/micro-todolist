package script

import (
	"context"
	"todolist/app/taskkx/repository/mq/task"
)

// 监听
func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskService(ctx)
	if err != nil {
		return
	}

}
