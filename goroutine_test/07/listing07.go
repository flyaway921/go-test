//这个示例程序展示goroutine调度器是如何在单个线程上切分时间片的
package main

import (
	"fmt"
	"runtime"
	"sync"
)

//wg 用来等待程序完成，相当于一个计数信号量
var wg sync.WaitGroup

func main() {
	//分配2个逻辑处理器给调度器使用，默认情况下该值应该是CPU核数
	runtime.GOMAXPROCS(2)
	//计数加2，表示要等待两个goroutine
	wg.Add(2)

	//创建两个goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	//等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

//显示5000以内的素数
func printPrime(prefix string) {
	defer wg.Done()

next:

	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
