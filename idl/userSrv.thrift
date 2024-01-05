namespace go api

struct UserModel{
    1:i64 id
    2:string username
    3:i64 created_at
    4:i64 updated_at
    5:i64 deleted_at
}

struct UserRequest{
    1:string username
    2:string password    3:string password_confirm
}

struct UserResponse{
    1:UserModel user
    2:i64 code
}


service MyService{
    UserResponse UserRegister(1:UserRequest req)
    UserResponse UserLogin(1:UserRequest req)
}
