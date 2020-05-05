package until_all_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func allResponse() string {
	numOfRunner := 10
	//ch := make(chan string) //使用chan 会导致僵尸协程，服务器资源耗尽
	ch := make(chan string, numOfRunner) // 使用buffered chan
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for i := 0; i < numOfRunner; i++ {
		finalRet += <-ch + "\n" //取出channel内数据并拼装
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("before:", runtime.NumGoroutine()) //获取运行时的协程数量
	t.Log(allResponse())
	time.Sleep(time.Second * 1)
	t.Log("after:", runtime.NumGoroutine()) //获取运行后的协程数量
}
