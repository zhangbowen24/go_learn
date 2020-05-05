package csp

import (
	"fmt"
	"testing"
	"time"
)

type msgData map[string]string

func service() msgData {
	time.Sleep(time.Millisecond * 50)
	res := msgData{"data": "done", "msg": "ok"}
	return res
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan msgData {
	//retCh := make(chan msgData) //channel
	retCh := make(chan msgData, 1) //buffer channel
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	ret := <-retCh
	fmt.Println(ret["data"])
	fmt.Println(ret["msg"])
}
