package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

//传递给channel空值
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

//关闭channel
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "canceled")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan)
	cancel_2(cancelChan)
}
