package rpc

import "C"
import (
	"context"
	"log"
	"todolist/idl/kitex_gen/api"
)

//func UserLogin(ctx context.Context, req *api.UserRequest) (resp *api.UserResponse, err error) {
//	for {
//		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
//		resp, err = Client.UserLogin(ctx, req)
//		cancel()
//		if err != nil {
//			log.Fatal(err)
//		}
//		log.Println(resp)
//		time.Sleep(time.Second)
//	}
//}

func UserRegister(ctx context.Context, req *api.UserRequest) (resp *api.UserResponse, err error) {
	if req != nil {
		//_, cancel := context.WithTimeout(context.Background(), time.Second*3)
		resp, err = Client.UserRegister(ctx, req)
		//cancel()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		//time.Sleep(time.Second)
	}
	return
}
