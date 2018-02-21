//这个示例程序展示如何在程序里造成竞争状态
//实际上不希望出现这种情况

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter是所有goroutine都要增加其值的变量
	counter int

	//wg是用来等待程序结束
	wg sync.WaitGroup
)

func main() {

	wg.Add(2)
	//创建两个goroutine
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//捕获counter的值
		value := counter
		//当前goroutine从线程退出，回到队列，等待被调度
		runtime.Gosched()
		//增加本地value变量的值
		value++
		//将该值保存回counter
		counter = value
	}
}
