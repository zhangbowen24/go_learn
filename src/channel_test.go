package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 7; i++ {
			ch <- i
		}
		close(ch)
		//ch <- 11 //往关闭的channel里发送数据会引发panic错误
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		/*for {
			if data, ok := <-ch; ok { // ok=false 代表通道关闭
				fmt.Println(data)
			} else {
				break
			}
		}*/
		for i := 0; i < 11; i++ {
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	/*wg.Add(1)
	dataReceiver(ch, &wg)*/
	wg.Wait()
}
