package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//值传递
		/*go func(i int) {
			fmt.Println(i)
		}(i)*/
		//共享关系
		/*go func() {
			fmt.Println(i)
		}()*/
	}
	time.Sleep(time.Millisecond * 50)
}
