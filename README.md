# threadpool

基于go实现的线程池，来完成百万级高并发

# 简介

TODO

# 安装

```bash
go get -u github.com/lk668/threadpool
```

# 使用

```go
package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/lk668/threadpool"
)

//需要执行任务的结构体
type Task struct {
	Number int
}

// 实现Job这个interface的RunTask函数
func (t Task) RunTask(request interface{}) {
	fmt.Println("This is task: ", t.Number)
	//设置个等待时间
	time.Sleep(1 * time.Second)
}

func main() {

	// 设置线程池的大小
	poolNum := 100 * 100 * 20
    jobQueueNum := 100
	workerPool := threadpool.GetWorkerPool(poolNum, jobQueueNum)
	workerPool.Start()

	// 模拟百万请求
	dataNum := 100 * 100 * 100

	go func() {
		for i := 0; i < dataNum; i++ {
			task := Task{Number: i}
			workerPool.JobQueue <- task
		}
	}()

	// 阻塞主线程
	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
```

