package threadpool

// 需要执行的job
type Job interface {
	RunTask(request interface{})
}

// Job channel
type JobChan chan Job
