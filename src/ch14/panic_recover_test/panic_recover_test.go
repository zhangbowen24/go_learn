package panic_recover_test

import (
	"errors"
	"fmt"
	//"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("finally")
	}()

	//recover 恢复机制，小心成为僵尸进程，如果真的出错了，就出错吧，如果可以重新启动进程就重新启动
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered form", err)
		}
	}()

	fmt.Println("Start")
	//fmt.Println("start")
	panic(errors.New("Something wrong"))
	//os.Exit(-1)
	//fmt.Println("start")
}
