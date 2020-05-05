package select_test

import (
	"fmt"
	"testing"
	"time"
)

type msgData map[string]string

func service() msgData {
	time.Sleep(time.Millisecond * 500)
	res := msgData{"data": "done", "msg": "ok"}
	return res
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
	//retCh := AsyncService()
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100): //程序超时判断
		t.Error("time out")
	}
}
