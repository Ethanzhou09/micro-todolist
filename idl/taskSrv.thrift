namespace go api

struct TaskModel{
    1:i64 id
    2:i64 uid
    3:string title
    4:string content
    5:i64 start_time
    6:i64 end_time
    7:i64 status
    8:i64 create_time
    9:i64 update_time
}

struct TaskRequest{
    1:i64 id
    2:i64 uid
    3:string title
    4:string content
    5:i64 start_time
    6:i64 end_time
    7:i64 status
    8:i64 start
    9:i64 limit
}
struct TaskListResponse{
    1:list<TaskModel> task_list
    2:i64 count
    3:i64 code
}

struct TaskDetailResponse{
    1:TaskModel task_detail
    2:i64 code
}

service TaskService{
    TaskListResponse GetTaskList(1:TaskRequest req)
    TaskDetailResponse GetTask(1:TaskRequest req)
    TaskDetailResponse CreateTask(1:TaskRequest req)
    TaskDetailResponse UpdateTask(1:TaskRequest req)
    TaskDetailResponse DeleteTask(1:TaskRequest req)
}