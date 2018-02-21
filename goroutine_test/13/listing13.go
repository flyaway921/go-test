//用原子操作来解决共享资源竞争问题

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter是所有goroutine都要增加其值的变量
	counter int64

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
		//用原子的AddInt64实现加，同样也可以做减法，乘法，除法
		atomic.AddInt64(&counter, 1)
		//当前goroutine从线程退出，回到队列，等待被调度
		runtime.Gosched()
		//增加本地value变量的值
	}
}
